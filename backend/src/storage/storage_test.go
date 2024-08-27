package storage

import (
	"database/sql"
	"log"
	"os"

	"search_scraper/src/types"
	"testing"

	_ "github.com/mattn/go-sqlite3"
	"github.com/stretchr/testify/assert"
)

func SetUpTest(t *testing.T) func(t *testing.T) {
	log.Println("setup test")

	return func(t *testing.T) {
		log.Println("teardown test")
	}
}
func SetUpTestDB(t *testing.T) (func(t *testing.T), *sql.DB) {
	db, err := sql.Open("sqlite3", "./db_test.sqlite")
	if err != nil {
		t.Fatal(err)
		return nil, nil
	}
	_, err = db.Exec("CREATE TABLE IF NOT EXISTS links (id INTEGER PRIMARY KEY, domain TEXT, url TEXT, filter_type TEXT)")
	if err != nil {
		t.Fatal(err)
	}

	return func(t *testing.T) {
		db.Close()
		log.Println("teardown test db")
		err := os.Remove("./db_test.sqlite")
		if err != nil {
			t.Fatal(err)
		}
	}, db
}

func TestInit(t *testing.T) {
	teardown, db := SetUpTestDB(t)
	defer teardown(t)
	type args struct {
		db *sql.DB
	}
	tests := map[string]struct {
		args
		want    *Storage
		wantErr error
	}{
		"db_is_nil": {
			args: args{
				db: nil,
			},
			want:    nil,
			wantErr: ErrDBIsNil,
		},
		"db_is_not_nil": {
			args: args{
				db: &sql.DB{},
			},
			want: &Storage{
				db: &sql.DB{},
			},
			wantErr: nil,
		},
		"db_is_real_db": {
			args: args{
				db: db,
			},
			want: &Storage{
				db: db,
			},
		},
	}

	for key, tt := range tests {
		t.Run(key, func(t *testing.T) {
			got, err := Init(tt.args.db)
			assert.Equal(t, tt.wantErr, err)
			assert.Equal(t, tt.want, got)
		})
	}
}

func TestGetList(t *testing.T) {
	tearDown, db := SetUpTestDB(t)
	defer tearDown(t)

	st, err := Init(db)
	if err != nil {
		t.Fatal(err)
	}
	type args struct {
		st       *Storage
		listType string
	}
	tests := map[string]struct {
		args
		setUp   func(t *testing.T) func(t *testing.T)
		want    []types.Link
		wantErr error
	}{
		"empty_database": {
			args: args{
				st:       st,
				listType: "links",
			},
			setUp: func(t *testing.T) func(t *testing.T) {
				return func(t *testing.T) {
					_, err = db.Exec("DELETE FROM links")
					if err != nil {
						t.Fatal(err)
					}
				}
			},
			want:    nil,
			wantErr: nil,
		},
		"valid_list_type_with_one_link": {
			args: args{
				st:       st,
				listType: "links",
			},
			setUp: func(t *testing.T) func(t *testing.T) {
				_, err = db.Exec("INSERT INTO links (domain, url, filter_type) VALUES ('example.com', 'http://example.com', 'domain')")
				if err != nil {
					t.Fatal(err)
				}
				return func(t *testing.T) {
					_, err = db.Exec("DELETE FROM links")

					if err != nil {
						t.Fatal(err)
					}
				}
			},
			want: []types.Link{
				{ID: 1, Domain: "example.com", Url: "http://example.com", FilterType: "domain"},
			},
			wantErr: nil,
		},
		"valid_list_type_with_multiple_links": {
			args: args{
				st:       st,
				listType: "links",
			},
			setUp: func(t *testing.T) func(t *testing.T) {
				_, err = db.Exec(`
				INSERT INTO links (domain, url, filter_type) VALUES ('example1.com', 'http://example1.com', 'domain');
				INSERT INTO links (domain, url, filter_type) VALUES ('example2.com', 'http://example2.com', 'domain');
				INSERT INTO links (domain, url, filter_type) VALUES ('example3.com', 'http://example3.com', 'domain');
				`)
				if err != nil {
					t.Fatal(err)
				}
				return func(t *testing.T) {
					_, err = db.Exec("DELETE FROM links")

					if err != nil {
						t.Fatal(err)
					}
				}
			},
			want: []types.Link{
				{ID: 1, Domain: "example1.com", Url: "http://example1.com", FilterType: "domain"},
				{ID: 2, Domain: "example2.com", Url: "http://example2.com", FilterType: "domain"},
				{ID: 3, Domain: "example3.com", Url: "http://example3.com", FilterType: "domain"},
			},
			wantErr: nil,
		},
	}

	for key, tt := range tests {
		t.Run(key, func(t *testing.T) {
			tearDown := tt.setUp(t)
			defer tearDown(t)

			got, err := tt.args.st.GetList(tt.args.listType)
			assert.Equal(t, tt.wantErr, err)
			assert.Equal(t, tt.want, got)
		})
	}
}

func TestAddLinkToList(t *testing.T) {
	tearDown, db := SetUpTestDB(t)
	defer tearDown(t)

	st, err := Init(db)
	if err != nil {
		t.Fatal(err)
	}
	type args struct {
		st       *Storage
		listType string
		link     types.Link
	}

	tests := map[string]struct {
		args
		setUp          func(t *testing.T) func(t *testing.T)
		wantDBInstance []types.Link
		want           types.Link
		wantErr        error
	}{
		"valid_link": {
			args: args{
				st:       st,
				listType: "links",
				link: types.Link{
					Domain:     "example.com",
					Url:        "http://example.com",
					FilterType: "domain",
				},
			},
			setUp: func(t *testing.T) func(t *testing.T) {
				_, err = db.Exec("DELETE FROM links")
				if err != nil {
					t.Fatal(err)
				}
				return func(t *testing.T) {
					_, err = db.Exec("DELETE FROM links")
					if err != nil {
						t.Fatal(err)
					}
				}
			},
			wantDBInstance: []types.Link{
				{ID: 1, Domain: "example.com", Url: "http://example.com", FilterType: "domain"},
			},
			want:    types.Link{ID: 1, Domain: "example.com", Url: "http://example.com", FilterType: "domain"},
			wantErr: nil,
		},
		"invalid_link_no_domain": {
			args: args{
				st:       st,
				listType: "links",
				link: types.Link{
					Domain:     "",
					Url:        "http://example.com",
					FilterType: "domain",
				},
			},
			setUp: func(t *testing.T) func(t *testing.T) {
				_, err = db.Exec("DELETE FROM links")
				if err != nil {
					t.Fatal(err)
				}
				return func(t *testing.T) {
					_, err = db.Exec("DELETE FROM links")
					if err != nil {
						t.Fatal(err)
					}
				}
			},
			wantDBInstance: []types.Link{
				{ID: 1, Domain: "example.com", Url: "http://example.com", FilterType: "domain"},
			},
			want:    types.Link{ID: 1, Domain: "example.com", Url: "http://example.com", FilterType: "domain"},
			wantErr: nil,
		},
		"invalid_link_no_url": {
			args: args{
				st:       st,
				listType: "links",
				link: types.Link{
					Domain:     "example.com",
					Url:        "",
					FilterType: "domain",
				},
			},
			setUp: func(t *testing.T) func(t *testing.T) {
				_, err = db.Exec("DELETE FROM links")
				if err != nil {
					t.Fatal(err)
				}
				return func(t *testing.T) {
					_, err = db.Exec("DELETE FROM links")
					if err != nil {
						t.Fatal(err)
					}
				}
			},
			wantDBInstance: nil,
			want:           types.Link{},
			wantErr:        ErrOnURL,
		},
		"invalid_link_no_filter_type": {
			args: args{
				st:       st,
				listType: "links",
				link: types.Link{
					Domain:     "example.com",
					Url:        "http://example.com",
					FilterType: "",
				},
			},
			setUp: func(t *testing.T) func(t *testing.T) {
				_, err = db.Exec("DELETE FROM links")
				if err != nil {
					t.Fatal(err)
				}
				return func(t *testing.T) {
					_, err = db.Exec("DELETE FROM links")
					if err != nil {
						t.Fatal(err)
					}
				}
			},
			wantDBInstance: []types.Link{
				{ID: 1, Domain: "example.com", Url: "http://example.com", FilterType: "domain"},
			},
			want:    types.Link{ID: 1, Domain: "example.com", Url: "http://example.com", FilterType: "domain"},
			wantErr: nil,
		},
	}

	for key, tt := range tests {
		t.Run(key, func(t *testing.T) {
			tearDown := tt.setUp(t)
			defer tearDown(t)

			link, err := st.AddLinkToList(tt.listType, tt.link)
			assert.Equal(t, err, tt.wantErr)
			assert.Equal(t, tt.want, link)

			links, err := st.GetList(tt.listType)
			if err != nil {
				t.Fatal(err)
			}
			assert.Equal(t, tt.wantDBInstance, links)
		})
	}
}

func TestGetLinkFromList(t *testing.T) {
	tearDown, db := SetUpTestDB(t)
	defer tearDown(t)
	st, err := Init(db)
	if err != nil {
		t.Fatal(err)
	}
	type args struct {
		st       *Storage
		listType string
		linkID   int
	}
	tests := map[string]struct {
		args
		setUp   func(t *testing.T) func(t *testing.T)
		want    types.Link
		wantErr error
	}{
		"empty_db": {
			args: args{
				st:       st,
				listType: "links",
				linkID:   1,
			},
			setUp: func(t *testing.T) func(t *testing.T) {
				_, err = db.Exec("DELETE FROM links")

				if err != nil {
					t.Fatal(err)
				}
				return func(t *testing.T) {
					_, err = db.Exec("DELETE FROM links")

					if err != nil {
						t.Fatal(err)
					}
				}
			},
			want:    types.Link{},
			wantErr: sql.ErrNoRows,
		},
		"valid_db_many_links_corect_id": {
			args: args{
				st:       st,
				listType: "links",
				linkID:   1,
			},
			setUp: func(t *testing.T) func(t *testing.T) {
				_, err = db.Exec(`
				INSERT INTO links (domain, url, filter_type) VALUES ('example1.com', 'http://example1.com', 'domain');
				INSERT INTO links (domain, url, filter_type) VALUES ('example2.com', 'http://example2.com', 'domain');
				INSERT INTO links (domain, url, filter_type) VALUES ('example3.com', 'http://example3.com', 'domain');
				`)
				if err != nil {
					t.Fatal(err)
				}
				return func(t *testing.T) {
					_, err = db.Exec("DELETE FROM links")

					if err != nil {
						t.Fatal(err)
					}
				}
			},
			want: types.Link{
				ID: 1, Domain: "example1.com", Url: "http://example1.com", FilterType: "domain",
			},
			wantErr: nil,
		},
		"valid_db_many_links_not_corect_id": {
			args: args{
				st:       st,
				listType: "links",
				linkID:   4,
			},
			setUp: func(t *testing.T) func(t *testing.T) {
				_, err = db.Exec(`
				INSERT INTO links (domain, url, filter_type) VALUES ('example1.com', 'http://example1.com', 'domain');
				INSERT INTO links (domain, url, filter_type) VALUES ('example2.com', 'http://example2.com', 'domain');
				INSERT INTO links (domain, url, filter_type) VALUES ('example3.com', 'http://example3.com', 'domain');
				`)
				if err != nil {
					t.Fatal(err)
				}
				return func(t *testing.T) {
					_, err = db.Exec("DELETE FROM links")

					if err != nil {
						t.Fatal(err)
					}
				}
			},
			want:    types.Link{},
			wantErr: sql.ErrNoRows,
		},
	}

	for key, tt := range tests {
		t.Run(key, func(t *testing.T) {
			tearDown := tt.setUp(t)
			defer tearDown(t)

			got, err := st.GetLinkFromList(tt.listType, tt.linkID)
			assert.Equal(t, err, tt.wantErr)
			assert.Equal(t, tt.want, got)
		})
	}
}

func TestUpdateLinkInList(t *testing.T) {
	tearDown, db := SetUpTestDB(t)
	defer tearDown(t)
	st, err := Init(db)
	if err != nil {
		t.Fatal(err)
	}
	type args struct {
		st       *Storage
		listType string
		link     types.Link
	}
	tests := map[string]struct {
		args
		setUp          func(t *testing.T) func(t *testing.T)
		wantDBInstance types.Link
		wantErr        error
	}{
		"empty_db_empty_link": {
			args: args{
				st:       st,
				listType: "links",
				link:     types.Link{},
			},
			setUp: func(t *testing.T) func(t *testing.T) {
				_, err = db.Exec("DELETE FROM links")
				if err != nil {
					t.Fatal(err)
				}
				return func(t *testing.T) {
					_, err = db.Exec("DELETE FROM links")
					if err != nil {
						t.Fatal(err)
					}
				}
			},
			wantErr:        sql.ErrNoRows,
			wantDBInstance: types.Link{},
		},
		"empty_db_filled_link": {
			args: args{
				st:       st,
				listType: "links",
				link: types.Link{
					ID:         1,
					Domain:     "example.com",
					Url:        "http://example.com",
					FilterType: "domain",
				},
			},
			setUp: func(t *testing.T) func(t *testing.T) {
				_, err = db.Exec("DELETE FROM links")
				if err != nil {
					t.Fatal(err)
				}
				return func(t *testing.T) {
					_, err = db.Exec("DELETE FROM links")
					if err != nil {
						t.Fatal(err)
					}
				}
			},
			wantErr:        sql.ErrNoRows,
			wantDBInstance: types.Link{},
		},
		"filled_db_link_only_with_valid_id": {
			args: args{
				st:       st,
				listType: "links",
				link: types.Link{
					ID: 1,
				},
			},
			setUp: func(t *testing.T) func(t *testing.T) {
				_, err = db.Exec("INSERT INTO links (domain, url, filter_type) VALUES ('example1.com', 'http://example1.com', 'domain');")
				if err != nil {
					t.Fatal(err)
				}
				return func(t *testing.T) {
					_, err = db.Exec("DELETE FROM links")
					if err != nil {
						t.Fatal(err)
					}
				}
			},
			wantErr: nil,
			wantDBInstance: types.Link{
				ID:         1,
				Url:        "http://example1.com",
				Domain:     "example1.com",
				FilterType: "domain",
			},
		},
		"filled_db_filled_link_valid_id": {
			args: args{
				st:       st,
				listType: "links",
				link: types.Link{
					ID:         1,
					Domain:     "example.com",
					Url:        "http://example.com",
					FilterType: "domain",
				},
			},
			setUp: func(t *testing.T) func(t *testing.T) {
				_, err = db.Exec("INSERT INTO links (domain, url, filter_type) VALUES ('example1.com', 'http://example1.com', 'domain');")
				if err != nil {
					t.Fatal(err)
				}
				return func(t *testing.T) {
					_, err = db.Exec("DELETE FROM links")
					if err != nil {
						t.Fatal(err)
					}
				}
			},
			wantErr: nil,
			wantDBInstance: types.Link{
				ID:         1,
				Domain:     "example.com",
				Url:        "http://example.com",
				FilterType: "domain",
			},
		},
		"filled_db_link_only_with_invalid_id": {
			args: args{
				st:       st,
				listType: "links",
				link: types.Link{
					ID: 2,
				},
			},
			setUp: func(t *testing.T) func(t *testing.T) {
				_, err = db.Exec("DELETE FROM links")
				_, err = db.Exec("INSERT INTO links (domain, url, filter_type) VALUES ('example1.com', 'http://example1.com', 'domain');")
				if err != nil {
					t.Fatal(err)
				}
				return func(t *testing.T) {
					_, err = db.Exec("DELETE FROM links")
					if err != nil {
						t.Fatal(err)
					}
				}
			},
			wantErr:        sql.ErrNoRows,
			wantDBInstance: types.Link{},
		},
		"filled_db_filled_link_invalid_id": {
			args: args{
				st:       st,
				listType: "links",
				link: types.Link{
					ID:         2,
					Domain:     "example.com",
					Url:        "http://example.com",
					FilterType: "domain",
				},
			},
			setUp: func(t *testing.T) func(t *testing.T) {
				_, err = db.Exec("DELETE FROM links")
				_, err = db.Exec("INSERT INTO links (domain, url, filter_type) VALUES ('example1.com', 'http://example1.com', 'domain');")
				if err != nil {
					t.Fatal(err)
				}
				return func(t *testing.T) {
					_, err = db.Exec("DELETE FROM links")
					if err != nil {
						t.Fatal(err)
					}
				}
			},
			wantErr:        sql.ErrNoRows,
			wantDBInstance: types.Link{},
		},
	}

	for key, tt := range tests {
		t.Run(key, func(t *testing.T) {
			tearDown := tt.setUp(t)
			defer tearDown(t)

			err := tt.args.st.UpdateLinkInList(tt.listType, tt.link)
			assert.Equal(t, tt.wantErr, err)

			got, _ := tt.args.st.GetLinkFromList(tt.listType, tt.link.ID)
			assert.Equal(t, tt.wantDBInstance, got)
		})
	}
}

func TestDeleteLinkFromList(t *testing.T) {
	tearDown, db := SetUpTestDB(t)
	defer tearDown(t)
	st, err := Init(db)
	if err != nil {
		t.Fatal(err)
	}
	type args struct {
		st       *Storage
		listType string
		linkID   int
	}
	tests := map[string]struct {
		args
		setUp   func(t *testing.T) func(t *testing.T)
		wantErr error
	}{
		"empty_db": {
			args: args{
				st:       st,
				listType: "links",
				linkID:   1,
			},
			setUp: func(t *testing.T) func(t *testing.T) {
				_, err = db.Exec("DELETE FROM links")
				if err != nil {
					t.Fatal(err)
				}
				return func(t *testing.T) {
					_, err = db.Exec("DELETE FROM links")
					if err != nil {
						t.Fatal(err)
					}
				}
			},
			wantErr: sql.ErrNoRows,
		},
		"filled_db_valid_id": {
			args: args{
				st:       st,
				listType: "links",
				linkID:   1,
			},
			setUp: func(t *testing.T) func(t *testing.T) {
				_, err = db.Exec("DELETE FROM links")
				_, err = db.Exec("INSERT INTO links (domain, url, filter_type) VALUES ('example1.com', 'http://example1.com', 'domain');")
				if err != nil {
					t.Fatal(err)
				}
				return func(t *testing.T) {
					_, err = db.Exec("DELETE FROM links")
					if err != nil {
						t.Fatal(err)
					}
				}
			},
			wantErr: nil,
		},
		"filled_db_invalid_id": {
			args: args{
				st:       st,
				listType: "links",
				linkID:   2,
			},
			setUp: func(t *testing.T) func(t *testing.T) {
				_, err = db.Exec("DELETE FROM links")
				_, err = db.Exec("INSERT INTO links (domain, url, filter_type) VALUES ('example1.com', 'http://example1.com', 'domain');")
				if err != nil {
					t.Fatal(err)
				}
				return func(t *testing.T) {
					_, err = db.Exec("DELETE FROM links")
					if err != nil {
						t.Fatal(err)
					}
				}
			},
			wantErr: sql.ErrNoRows,
		},
	}

	for key, tt := range tests {
		t.Run(key, func(t *testing.T) {
			tearDown := tt.setUp(t)
			defer tearDown(t)

			err := tt.args.st.DeleteLinkFromList(tt.listType, tt.linkID)
			assert.Equal(t, tt.wantErr, err)
		})
	}
}

func TestConteinsLinkIntoList(t *testing.T) {
	tearDown, db := SetUpTestDB(t)
	defer tearDown(t)

	st, err := Init(db)
	if err != nil {
		t.Fatal(err)
	}

	type args struct {
		st       *Storage
		listType string
		link     types.Link
	}

	tests := map[string]struct {
		args
		setUp   func(t *testing.T) func(t *testing.T)
		wantErr error
		want    bool
	}{
		"empty_database": {
			args: args{
				st:       st,
				listType: "links",
				link: types.Link{
					Domain: "example1.com",
					Url:    "http://example1.com",
				},
			},
			setUp: func(t *testing.T) func(t *testing.T) {
				return func(t *testing.T) {
					_, err = db.Exec("DELETE FROM links")
					if err != nil {
						t.Fatal(err)
					}
				}
			},
			wantErr: nil,
			want:    false,
		},
		"filled_database_conteins_link": {
			args: args{
				st:       st,
				listType: "links",
				link: types.Link{
					Domain: "example1.com",
					Url:    "http://example1.com",
				},
			},
			setUp: func(t *testing.T) func(t *testing.T) {
				_, err = db.Exec(`
				INSERT INTO links (domain, url, filter_type) VALUES ('example1.com', 'http://example1.com', 'domain');
				INSERT INTO links (domain, url, filter_type) VALUES ('example2.com', 'http://example2.com', 'domain');
				INSERT INTO links (domain, url, filter_type) VALUES ('example3.com', 'http://example3.com', 'domain');
				`)
				if err != nil {
					t.Fatal(err)
				}
				return func(t *testing.T) {
					_, err = db.Exec("DELETE FROM links")
					if err != nil {
						t.Fatal(err)
					}
				}
			},
			wantErr: nil,
			want:    true,
		},
		"filled_database_conteins_multiple_link": {
			args: args{
				st:       st,
				listType: "links",
				link: types.Link{
					Domain: "example3.com",
					Url:    "http://example3.com",
				},
			},
			setUp: func(t *testing.T) func(t *testing.T) {
				_, err = db.Exec(`
				INSERT INTO links (domain, url, filter_type) VALUES ('example1.com', 'http://example1.com', 'domain');
				INSERT INTO links (domain, url, filter_type) VALUES ('example3.com', 'http://example3.com', 'domain');
				INSERT INTO links (domain, url, filter_type) VALUES ('example3.com', 'http://example3.com', 'domain');
				`)
				if err != nil {
					t.Fatal(err)
				}
				return func(t *testing.T) {
					_, err = db.Exec("DELETE FROM links")
					if err != nil {
						t.Fatal(err)
					}
				}
			},
			wantErr: nil,
			want:    true,
		},
		"filled_database_dosent_contein_link": {
			args: args{
				st:       st,
				listType: "links",
				link: types.Link{
					Domain: "example4.com",
					Url:    "http://example4.com",
				},
			},
			setUp: func(t *testing.T) func(t *testing.T) {
				_, err = db.Exec(`
				INSERT INTO links (domain, url, filter_type) VALUES ('example1.com', 'http://example1.com', 'domain');
				INSERT INTO links (domain, url, filter_type) VALUES ('example2.com', 'http://example2.com', 'domain');
				INSERT INTO links (domain, url, filter_type) VALUES ('example3.com', 'http://example3.com', 'domain');
				`)
				if err != nil {
					t.Fatal(err)
				}
				return func(t *testing.T) {
					_, err = db.Exec("DELETE FROM links")
					if err != nil {
						t.Fatal(err)
					}
				}
			},
			wantErr: nil,
			want:    false,
		},
	}

	for key, tt := range tests {
		t.Run(key, func(t *testing.T) {
			tearDown := tt.setUp(t)
			defer tearDown(t)

			got, err := tt.args.st.ConteinsLinkInList(tt.listType, tt.link)
			assert.Equal(t, tt.wantErr, err)
			assert.Equal(t, tt.want, got)
		})
	}
}
