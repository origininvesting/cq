package test

import (
	"database/sql/driver"
	"flag"
	"github.com/wfreeman/cq"
	. "launchpad.net/gocheck"
	"log"
)

type ConnSuite struct{}

var (
	_       = Suite(&ConnSuite{})
	testURL = flag.String("testdb", "http://localhost:7474/", "the base url for the test db")
)

func openTest() driver.Conn {
	db, err := cq.Open(*testURL)
	if err != nil {
		log.Println("can't connect to db.")
		return nil
	}
	return db
}

func (s *ConnSuite) TestOpen(c *C) {
	db := openTest()
	if db == nil {
		c.Fatal("can't connect to test db: ", *testURL)
	}
}

func (s *ConnSuite) TestPrepareNoParams(c *C) {
	db := openTest()
	if db == nil {
		c.Fatal("can't connect to test db: ", *testURL)
	}
	stmt, err := db.Prepare("match (n) return n limit 1")
	c.Assert(err, IsNil)
	if stmt == nil {
		c.Fatal("stmt is nil")
	}
}

func (s *ConnSuite) TestBadURL(c *C) {
	db, err := cq.Open("")
	if err == nil {
		c.Fatal("err was nil!")
	}
	if db != nil {
		c.Fatal("db should be nil:", db)
	}
}