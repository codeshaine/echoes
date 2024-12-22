package store

import (
	"context"
	"database/sql"
)

type EchoData struct {
	ID            int    `json:"id"`
	Echoe         string `json:"echoe"`
	Author        string `josn:"author"`
	PublishedDate string `json:"published_date"`
}

type EchoWithLikes struct {
	Data  EchoData `json:"data"`
	Likes int      `json:"likes"`
}

type EchoList []EchoWithLikes

type EchoesStore struct {
	db *sql.DB
}

func (e *EchoesStore) GetAll(ctx context.Context) (EchoList, error) {
	// get all echoes
	const query = `SELECT e.id, e.echo, e.author, e.created_at, COUNT(l.id)
		FROM echoes AS e
    	FULL JOIN likes As l ON e.id = l.echo_id
		GROUP BY e.id
		ORDER BY e.created_at DESC;
	`
	c, cancel := context.WithTimeout(ctx, QueryTimeout)
	defer cancel()
	res, err := e.db.QueryContext(c, query)
	if err != nil {
		return nil, err
	}
	defer res.Close()
	//how to scan the row and fill it into the EchoList
	var echoes EchoList
	for res.Next() {
		var echo EchoWithLikes
		err := res.Scan(&echo.Data.ID, &echo.Data.Echoe, &echo.Data.Author, &echo.Data.PublishedDate, &echo.Likes)
		if err != nil {
			return nil, err
		}
		echoes = append(echoes, echo)
	}
	return echoes, nil

}

func (e *EchoesStore) GetPopular() {
	// get popular echoes
}
func (e *EchoesStore) GetTrending() {
	// get trending echoes
}
