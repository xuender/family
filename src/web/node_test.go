package web

import (
	log "github.com/Sirupsen/logrus"
	"gopkg.in/mgo.v2/bson"
	"testing"
)

func TestNodeFind(t *testing.T) {
	s := Connect()
	defer s.Close()
	DB = GetTestDb(s)
	DB.DropDatabase()
	c := DB.C("node")
	n := NodeNew(Data{})
	err := c.Insert(&n)
	if err != nil {
		log.Error(err)
		t.Errorf("增加Node失败")
	}
	_, e := NodeFind(n.Id)
	if e != nil {
		log.Error(e)
		t.Errorf("查找Node失败")
	}
}
func TestNodeNew(t *testing.T) {
	n := NodeNew(Data{N: "test"})
	if n.N != "test" {
		t.Errorf("新建Node失败")
	}
}
func TestNodeSave(t *testing.T) {
	s := Connect()
	defer s.Close()
	DB = GetTestDb(s)
	DB.DropDatabase()
	n := NodeNew(Data{N: "test"})
	n.Save(bson.NewObjectId())
	if n.Ca.IsZero() || !n.Ua.IsZero() {
		t.Errorf("新建保存失败")
	}
	n.Save(bson.NewObjectId())
	if n.Ca.IsZero() || n.Ua.IsZero() {
		t.Errorf("新建再次保存失败")
	}
}
func TestNodeAdd(t *testing.T) {
	s := Connect()
	defer s.Close()
	DB = GetTestDb(s)
	DB.DropDatabase()
	n := NodeNew(Data{N: "test"})
	n.Save(bson.NewObjectId())
	if n.Ca.IsZero() || !n.Ua.IsZero() {
		t.Errorf("新建保存失败")
	}
	f := n.Add("f")
	if f.Ca.IsZero() || f.N != "父亲的姓名" {
		t.Errorf("增加父亲失败")
	}
	nf := n.Add("f")
	if f.N != "父亲的姓名" || f.Id != nf.Id {
		t.Errorf("增加父亲失败")
	}
	m := n.Add("m")
	if m.Ca.IsZero() || m.N != "母亲的姓名" {
		t.Errorf("增加母亲失败")
	}
	nm := n.Add("m")
	if m.N != "母亲的姓名" || m.Id != nm.Id {
		t.Errorf("增加母亲失败")
	}
	p := n.Add("p")
	if p.Ca.IsZero() || p.N != "丈夫的姓名" {
		t.Errorf("增加丈夫失败")
	}
	np := n.Add("p")
	if p.N != "丈夫的姓名" || p.Id != np.Id {
		t.Errorf("增加丈夫失败")
	}
}
func TestNodeRoot(t *testing.T) {
	s := Connect()
	defer s.Close()
	DB = GetTestDb(s)
	DB.DropDatabase()
	n := NodeNew(Data{N: "test"})
	n.Save(bson.NewObjectId())
	f := n.Add("f")
	g := f.Add("f")
	root := n.Root(4)
	if root.Id != g.Id {
		t.Errorf("根节点查找错误")
	}
}
func TestNodePartner(t *testing.T) {
	s := Connect()
	defer s.Close()
	DB = GetTestDb(s)
	DB.DropDatabase()
	n := NodeNew(Data{N: "test"})
	n.Save(bson.NewObjectId())
	p := n.Add("p")
	fp, err := n.Partner()
	if err != nil || p.Id != fp.Id {
		t.Errorf("伴侣查找错误")
	}
}
func TestNodeChildren(t *testing.T) {
	s := Connect()
	defer s.Close()
	DB = GetTestDb(s)
	DB.DropDatabase()
	n := NodeNew(Data{N: "test"})
	n.Save(bson.NewObjectId())
	f := n.Add("f")
	n2 := NodeNew(Data{N: "test"})
	n2.F = f.Id
	n2.Save(bson.NewObjectId())
	ns, err := f.Children()
	if err != nil || len(ns) != 2 {
		t.Errorf("子女查找错误")
	}
}
func TestNodeAddC(t *testing.T) {
	s := Connect()
	defer s.Close()
	DB = GetTestDb(s)
	DB.DropDatabase()
	n := NodeNew(Data{N: "test"})
	n.Save(bson.NewObjectId())
	f := n.Add("f")
	f.Add("c")
	ns, err := f.Children()
	if err != nil || len(ns) != 2 {
		t.Errorf("增加子女错误")
	}
}