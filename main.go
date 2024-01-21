package main

import (
	"fmt"
	"strings"
)

type Filter struct {
	ID       int64  // provee el id a buscar
	Name     string // provee el nombre a buscar
	Lastname string
	Search   string // provee un valor por cual buscar usando LIKE a nombres o apellidos
	Limit    int64
	Offset   int64
}

const qb = `SELECT * FROM movimientos WHERE 1=1 %ATTRS% %ID%  %FIRST_NAME% %LAST_NAME%  %LIMIT% %OFFSET%`

func QueryBuilderEvil(f Filter) (string, []any) {
	sql := qb
	var params []any

	if strings.TrimSpace(f.Search) != "" {
		sql = strings.ReplaceAll(sql, "%ATTRS%", " AND ( name LIKE ? OR last_name LIKE ?)")
		params = append(params, "%"+f.Search+"%", "%"+f.Search+"%")
	} else {
		sql = strings.ReplaceAll(sql, "%ATTRS%", "")
	}

	if f.ID > 0 {
		sql = strings.ReplaceAll(sql, "%ID%", " AND id = ?")
		params = append(params, f.ID)
	} else {
		sql = strings.ReplaceAll(sql, "%ID%", "")
	}

	if strings.TrimSpace(f.Name) != "" {
		sql = strings.ReplaceAll(sql, "%FIRST_NAME%", " AND name = ?")
		params = append(params, f.Name)
	} else {
		sql = strings.ReplaceAll(sql, "%FIRST_NAME%", "")
	}

	if strings.TrimSpace(f.Lastname) != "" {
		sql = strings.ReplaceAll(sql, "%LAST_NAME%", " AND last_name = ?")
		params = append(params, f.Lastname)
	} else {
		sql = strings.ReplaceAll(sql, "%LAST_NAME%", "")
	}

	if f.Limit > 0 {
		sql = strings.ReplaceAll(sql, "%LIMIT%", " LIMIT ?")
		params = append(params, f.Limit)
	} else {
		sql = strings.ReplaceAll(sql, "%LIMIT%", "")
	}

	if f.Offset > 0 {
		sql = strings.ReplaceAll(sql, "%OFFSET%", " OFFSET ?")
		params = append(params, f.Offset)
	} else {
		sql = strings.ReplaceAll(sql, "%OFFSET%", "")
	}

	return sql, params
}

const q = `SELECT id, name, last_name FROM clientes WHERE 1=1`

func QueryBuilderOK(f Filter) (strings.Builder, []any) {
	b := strings.Builder{}
	params := make([]any, 0, 7) // 7 es la cantidad máxima de elementos que se pueden agregar al filtro

	b.WriteString(q)

	if f.ID > 0 {
		b.WriteString(" AND id = ?") // Agregamos el segmento del filtro sql y el parámetro posicional
		params = append(params, f.ID)
	}

	if strings.TrimSpace(f.Name) != "" {
		//             Notese el espacio extra antes del sql
		b.WriteString(" AND name = ?")
		params = append(params, f.Name)
	}

	if strings.TrimSpace(f.Lastname) != "" {
		b.WriteString(" AND last_name = ?")
		params = append(params, f.Lastname)
	}

	if strings.TrimSpace(f.Search) != "" {
		b.WriteString(" AND ( name LIKE ? OR last_name LIKE ?)")
		params = append(params, fmt.Sprintf("%%%s%%", f.Search), fmt.Sprintf("%%%s%%", f.Search))
	}

	if f.Limit > 0 {
		b.WriteString(" LIMIT ?")
		params = append(params, f.Limit)
	}

	if f.Offset > 0 {
		b.WriteString(" OFFSET ?")
		params = append(params, f.Offset)
	}

	return b, params
}

func QueryBuilderOKAlter(f Filter) (strings.Builder, []any) {
	b := strings.Builder{}
	params := make([]any, 0, 7) // 7 es la cantidad máxima de elementos que se pueden agregar al filtro

	b.WriteString(q)

	if f.ID > 0 {
		b.WriteString(" AND id = ?") // Agregamos el segmento del filtro sql y el parámetro posicional
		params = append(params, f.ID)
	}

	if strings.TrimSpace(f.Name) != "" {
		//             Notese el espacio extra antes del sql
		b.WriteString(" AND name = ?")
		params = append(params, f.Name)
	}

	if strings.TrimSpace(f.Lastname) != "" {
		b.WriteString(" AND last_name = ?")
		params = append(params, f.Lastname)
	}

	if strings.TrimSpace(f.Search) != "" {
		b.WriteString(" AND ( name LIKE ? OR last_name LIKE ?)")
		b2 := strings.Builder{}
		b2.WriteString("%")
		b2.WriteString(f.Search)
		b2.WriteString("%")
		params = append(params, b2.String())
		b2.Reset()
		b2.WriteString("%")
		b2.WriteString(f.Search)
		b2.WriteString("%")
		params = append(params, b2.String())
	}

	if f.Limit > 0 {
		b.WriteString(" LIMIT ?")
		params = append(params, f.Limit)
	}

	if f.Offset > 0 {
		b.WriteString(" OFFSET ?")
		params = append(params, f.Offset)
	}

	return b, params
}

func QueryBuilderOKMask(f Filter) (strings.Builder, []any) {
	b := strings.Builder{}
	mask := struct {
		ID       bool
		Name     bool
		Lastname bool
		Search   bool
		Limit    bool
		Offset   bool
	}{}

	b.WriteString(q)

	size := 0

	if f.ID > 0 {
		b.WriteString(" AND id = ?") // Agregamos el segmento del filtro sql y el parámetro posicional
		mask.ID = true
		size++
	}

	if strings.TrimSpace(f.Name) != "" {
		b.WriteString(" AND name = ?")
		mask.Name = true
		size++
	}

	if strings.TrimSpace(f.Lastname) != "" {
		b.WriteString(" AND last_name = ?")
		mask.Lastname = true
		size++
	}

	if strings.TrimSpace(f.Search) != "" {
		b.WriteString(" AND ( name LIKE ? OR last_name LIKE ?)")
		mask.Search = true
		size += 2
	}

	if f.Limit > 0 {
		b.WriteString(" LIMIT ?")
		mask.Limit = true
		size++
	}

	if f.Offset > 0 {
		b.WriteString(" OFFSET ?")
		mask.Offset = true
		size++
	}

	p := make([]any, size)
	i := 0

	if mask.ID {
		p[i] = f.ID
		i++
	}

	if mask.Name {
		p[i] = f.Name
		i++
	}

	if mask.Lastname {
		p[i] = f.Lastname
		i++
	}

	if mask.Search {
		p[i] = fmt.Sprintf("%%%s%%", f.Search)
		i++
		p[i] = fmt.Sprintf("%%%s%%", f.Search)
		i++
	}

	if mask.Limit {
		p[i] = f.Limit
		i++
	}

	if mask.Offset {
		p[i] = f.Offset
	}

	return b, p
}
