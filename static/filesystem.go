// Code generated by shogo82148/assets-life v1.0.0. DO NOT EDIT.

//go:generate go run assets-life.go "." . static

package static

import (
	"io"
	"net/http"
	"os"
	"path"
	"sort"
	"strings"
	"time"
)

// Root is the root of the file system.
var Root http.FileSystem = fileSystem{
	file{
		name:    "/",
		content: "",
		mode:    0755 | os.ModeDir,
		next:    0,
		child:   1,
	},
	file{
		name:    "/filesystem.go",
		content: "// Code generated by shogo82148/assets-life v1.0.0. DO NOT EDIT.\n\n//go:generate go run assets-life.go \".\" . static\n\npackage static\n\nimport (\n\t\"io\"\n\t\"net/http\"\n\t\"os\"\n\t\"path\"\n\t\"sort\"\n\t\"strings\"\n\t\"time\"\n)\n\n// Root is the root of the file system.\nvar Root http.FileSystem = fileSystem{\n\tfile{\n\t\tname:    \"/\",\n\t\tcontent: \"\",\n\t\tmode:    0755 | os.ModeDir,\n\t\tnext:    0,\n\t\tchild:   1,\n\t},\n\tfile{\n\t\tname:    \"/filesystem.go\",\n",
		mode:    0644,
		next:    2,
		child:   -1,
	},
	file{
		name:    "/generate_tbl_ddl.sql",
		content: "SELECT\n  table_id,\n  REGEXP_REPLACE (schemaname, '^zzzzzzzz', '') AS schemaname,\n  REGEXP_REPLACE (tablename, '^zzzzzzzz', '') AS tablename,\n  seq,\n  ddl\nFROM\n  (\n    SELECT\n      table_id,\n      schemaname,\n      tablename,\n      seq,\n      ddl\n    FROM\n      (\n        --CREATE TABLE\n        SELECT\n          c.oid :: bigint as table_id,\n          n.nspname AS schemaname,\n          c.relname AS tablename,\n          2 AS seq,\n          'CREATE TABLE IF NOT EXISTS ' + QUOTE_IDENT(n.nspname) + '.' + QUOTE_IDENT(c.relname) + '' AS ddl\n        FROM\n          pg_namespace AS n\n          INNER JOIN pg_class AS c ON n.oid = c.relnamespace\n        WHERE\n          c.relkind = 'r' --OPEN PAREN COLUMN LIST\n        UNION\n        SELECT\n          c.oid :: bigint as table_id,\n          n.nspname AS schemaname,\n          c.relname AS tablename,\n          5 AS seq,\n          '(' AS ddl\n        FROM\n          pg_namespace AS n\n          INNER JOIN pg_class AS c ON n.oid = c.relnamespace\n        WHERE\n          c.relkind = 'r' --COLUMN LIST\n        UNION\n        SELECT\n          table_id,\n          schemaname,\n          tablename,\n          seq,\n          '\\t' + col_delim + col_name + ' ' + col_datatype + ' ' + col_nullable + ' ' + col_default + ' ' + col_encoding AS ddl\n        FROM\n          (\n            SELECT\n              c.oid :: bigint as table_id,\n              n.nspname AS schemaname,\n              c.relname AS tablename,\n              100000000 + a.attnum AS seq,\nCASE\n                WHEN a.attnum > 1 THEN ','\n                ELSE ''\n              END AS col_delim,\n              QUOTE_IDENT(a.attname) AS col_name,\nCASE\n                WHEN STRPOS(\n                  UPPER(format_type(a.atttypid, a.atttypmod)),\n                  'CHARACTER VARYING'\n                ) > 0 THEN REPLACE(\n                  UPPER(format_type(a.atttypid, a.atttypmod)),\n                  'CHARACTER VARYING',\n                  'VARCHAR'\n                )\n                WHEN STRPOS(\n                  UPPER(format_type(a.atttypid, a.atttypmod)),\n                  'CHARACTER'\n                ) > 0 THEN REPLACE(\n                  UPPER(format_type(a.atttypid, a.atttypmod)),\n                  'CHARACTER',\n                  'CHAR'\n                )\n                ELSE UPPER(format_type(a.atttypid, a.atttypmod))\n              END AS col_datatype,\nCASE\n                WHEN format_encoding((a.attencodingtype) :: integer) = 'none' THEN 'ENCODE RAW'\n                ELSE 'ENCODE ' + format_encoding((a.attencodingtype) :: integer)\n              END AS col_encoding,\nCASE\n                WHEN a.atthasdef IS TRUE THEN 'DEFAULT ' + adef.adsrc\n                ELSE ''\n              END AS col_default,\nCASE\n                WHEN a.attnotnull IS TRUE THEN 'NOT NULL'\n                ELSE ''\n              END AS col_nullable\n            FROM\n              pg_namespace AS n\n              INNER JOIN pg_class AS c ON n.oid = c.relnamespace\n              INNER JOIN pg_attribute AS a ON c.oid = a.attrelid\n              LEFT OUTER JOIN pg_attrdef AS adef ON a.attrelid = adef.adrelid\n              AND a.attnum = adef.adnum\n            WHERE\n              c.relkind = 'r'\n              AND a.attnum > 0\n            ORDER BY\n              a.attnum\n          ) --CONSTRAINT LIST\n        UNION\n        (\n          SELECT\n            c.oid :: bigint as table_id,\n            n.nspname AS schemaname,\n            c.relname AS tablename,\n            200000000 + CAST(con.oid AS INT) AS seq,\n            '\\t,' + pg_get_constraintdef(con.oid) AS ddl\n          FROM\n            pg_constraint AS con\n            INNER JOIN pg_class AS c ON c.relnamespace = con.connamespace\n            AND c.oid = con.conrelid\n            INNER JOIN pg_namespace AS n ON n.oid = c.relnamespace\n          WHERE\n            c.relkind = 'r'\n            AND pg_get_constraintdef(con.oid) NOT LIKE 'FOREIGN KEY%'\n          ORDER BY\n            seq\n        ) --CLOSE PAREN COLUMN LIST\n        UNION\n        SELECT\n          c.oid :: bigint as table_id,\n          n.nspname AS schemaname,\n          c.relname AS tablename,\n          299999999 AS seq,\n          ')' AS ddl\n        FROM\n          pg_namespace AS n\n          INNER JOIN pg_class AS c ON n.oid = c.relnamespace\n        WHERE\n          c.relkind = 'r' --BACKUP\n        UNION\n        SELECT\n          c.oid :: bigint as table_id,\n          n.nspname AS schemaname,\n          c.relname AS tablename,\n          300000000 AS seq,\n          'BACKUP NO' as ddl\n        FROM\n          pg_namespace AS n\n          INNER JOIN pg_class AS c ON n.oid = c.relnamespace\n          INNER JOIN (\n            SELECT\n              SPLIT_PART(key, '_', 5) id\n            FROM\n              pg_conf\n            WHERE\n              key LIKE 'pg_class_backup_%'\n              AND SPLIT_PART(key, '_', 4) = (\n                SELECT\n                  oid\n                FROM\n                  pg_database\n                WHERE\n                  datname = current_database()\n              )\n          ) t ON t.id = c.oid\n        WHERE\n          c.relkind = 'r' --BACKUP WARNING\n        UNION\n        SELECT\n          c.oid :: bigint as table_id,\n          n.nspname AS schemaname,\n          c.relname AS tablename,\n          1 AS seq,\n          '--WARNING: This DDL inherited the BACKUP NO property from the source table' as ddl\n        FROM\n          pg_namespace AS n\n          INNER JOIN pg_class AS c ON n.oid = c.relnamespace\n          INNER JOIN (\n            SELECT\n              SPLIT_PART(key, '_', 5) id\n            FROM\n              pg_conf\n            WHERE\n              key LIKE 'pg_class_backup_%'\n              AND SPLIT_PART(key, '_', 4) = (\n                SELECT\n                  oid\n                FROM\n                  pg_database\n                WHERE\n                  datname = current_database()\n              )\n          ) t ON t.id = c.oid\n        WHERE\n          c.relkind = 'r' --DISTSTYLE\n        UNION\n        SELECT\n          c.oid :: bigint as table_id,\n          n.nspname AS schemaname,\n          c.relname AS tablename,\n          300000001 AS seq,\nCASE\n            WHEN c.reldiststyle = 0 THEN 'DISTSTYLE EVEN'\n            WHEN c.reldiststyle = 1 THEN 'DISTSTYLE KEY'\n            WHEN c.reldiststyle = 8 THEN 'DISTSTYLE ALL'\n            WHEN c.reldiststyle = 9 THEN 'DISTSTYLE AUTO'\n            ELSE '<<Error - UNKNOWN DISTSTYLE>>'\n          END AS ddl\n        FROM\n          pg_namespace AS n\n          INNER JOIN pg_class AS c ON n.oid = c.relnamespace\n        WHERE\n          c.relkind = 'r' --DISTKEY COLUMNS\n        UNION\n        SELECT\n          c.oid :: bigint as table_id,\n          n.nspname AS schemaname,\n          c.relname AS tablename,\n          400000000 + a.attnum AS seq,\n          ' DISTKEY (' + QUOTE_IDENT(a.attname) + ')' AS ddl\n        FROM\n          pg_namespace AS n\n          INNER JOIN pg_class AS c ON n.oid = c.relnamespace\n          INNER JOIN pg_attribute AS a ON c.oid = a.attrelid\n        WHERE\n          c.relkind = 'r'\n          AND a.attisdistkey IS TRUE\n          AND a.attnum > 0 --SORTKEY COLUMNS\n        UNION\n        select\n          table_id,\n          schemaname,\n          tablename,\n          seq,\n          case\n            when min_sort < 0 then 'INTERLEAVED SORTKEY ('\n            else ' SORTKEY ('\n          end as ddl\n        from\n          (\n            SELECT\n              c.oid :: bigint as table_id,\n              n.nspname AS schemaname,\n              c.relname AS tablename,\n              499999999 AS seq,\n              min(attsortkeyord) min_sort\n            FROM\n              pg_namespace AS n\n              INNER JOIN pg_class AS c ON n.oid = c.relnamespace\n              INNER JOIN pg_attribute AS a ON c.oid = a.attrelid\n            WHERE\n              c.relkind = 'r'\n              AND abs(a.attsortkeyord) > 0\n              AND a.attnum > 0\n            group by\n              1,\n              2,\n              3,\n              4\n          )\n        UNION\n        (\n          SELECT\n            c.oid :: bigint as table_id,\n            n.nspname AS schemaname,\n            c.relname AS tablename,\n            500000000 + abs(a.attsortkeyord) AS seq,\nCASE\n              WHEN abs(a.attsortkeyord) = 1 THEN '\\t' + QUOTE_IDENT(a.attname)\n              ELSE '\\t, ' + QUOTE_IDENT(a.attname)\n            END AS ddl\n          FROM\n            pg_namespace AS n\n            INNER JOIN pg_class AS c ON n.oid = c.relnamespace\n            INNER JOIN pg_attribute AS a ON c.oid = a.attrelid\n          WHERE\n            c.relkind = 'r'\n            AND abs(a.attsortkeyord) > 0\n            AND a.attnum > 0\n          ORDER BY\n            abs(a.attsortkeyord)\n        )\n        UNION\n        SELECT\n          c.oid :: bigint as table_id,\n          n.nspname AS schemaname,\n          c.relname AS tablename,\n          599999999 AS seq,\n          '\\t)' AS ddl\n        FROM\n          pg_namespace AS n\n          INNER JOIN pg_class AS c ON n.oid = c.relnamespace\n          INNER JOIN pg_attribute AS a ON c.oid = a.attrelid\n        WHERE\n          c.relkind = 'r'\n          AND abs(a.attsortkeyord) > 0\n          AND a.attnum > 0 --END SEMICOLON\n        UNION\n        SELECT\n          c.oid :: bigint as table_id,\n          n.nspname AS schemaname,\n          c.relname AS tablename,\n          600000000 AS seq,\n          ';' AS ddl\n        FROM\n          pg_namespace AS n\n          INNER JOIN pg_class AS c ON n.oid = c.relnamespace\n        WHERE\n          c.relkind = 'r'\n      )\n    WHERE\n      tablename = $1\n    ORDER BY\n      table_id,\n      schemaname,\n      tablename,\n      seq\n  );\n",
		mode:    0644,
		next:    -1,
		child:   -1,
	},
}

type fileSystem []file

func (fs fileSystem) Open(name string) (http.File, error) {
	name = path.Clean("/" + name)
	i := sort.Search(len(fs), func(i int) bool { return fs[i].name >= name })
	if i >= len(fs) || fs[i].name != name {
		return nil, &os.PathError{
			Op:   "open",
			Path: name,
			Err:  os.ErrNotExist,
		}
	}
	f := &fs[i]
	return &httpFile{
		Reader: strings.NewReader(f.content),
		file:   f,
		fs:     fs,
		idx:    i,
		dirIdx: f.child,
	}, nil
}

type file struct {
	name    string
	content string
	mode    os.FileMode
	child   int
	next    int
}

var _ os.FileInfo = (*file)(nil)

func (f *file) Name() string {
	return path.Base(f.name)
}

func (f *file) Size() int64 {
	return int64(len(f.content))
}

func (f *file) Mode() os.FileMode {
	return f.mode
}

var zeroTime time.Time

func (f *file) ModTime() time.Time {
	return zeroTime
}

func (f *file) IsDir() bool {
	return f.Mode().IsDir()
}

func (f *file) Sys() interface{} {
	return nil
}

type httpFile struct {
	*strings.Reader
	file   *file
	fs     fileSystem
	idx    int
	dirIdx int
}

var _ http.File = (*httpFile)(nil)

func (f *httpFile) Stat() (os.FileInfo, error) {
	return f.file, nil
}

func (f *httpFile) Readdir(count int) ([]os.FileInfo, error) {
	ret := []os.FileInfo{}
	if !f.file.IsDir() {
		return ret, nil
	}

	if count <= 0 {
		for f.dirIdx >= 0 {
			entry := &f.fs[f.dirIdx]
			ret = append(ret, entry)
			f.dirIdx = entry.next
		}
		return ret, nil
	}

	ret = make([]os.FileInfo, 0, count)
	for f.dirIdx >= 0 {
		entry := &f.fs[f.dirIdx]
		ret = append(ret, entry)
		f.dirIdx = entry.next
		if len(ret) == count {
			return ret, nil
		}
	}
	return ret, io.EOF
}

func (f *httpFile) Close() error {
	return nil
}
