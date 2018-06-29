package main

import (
	"fmt"
	netURL "net/url"
	"rediscli/redis"
	redigo "github.com/gomodule/redigo/redis"
)

type RedisTest struct {
	nc redis.UnpooledConnection
	p  redis.Pool

}

func newConnectTest() *RedisTest {
	url := "redis://192.168.120.115:6379"
	parsedURL, _ := netURL.Parse(url)
	c, err := redis.NewConnection(parsedURL)
	if err != nil {
		panic(err)
	}

	p, err := redis.NewPool(url, redis.DefaultConfig)
	if err != nil {
		panic(err)
	}

	return &RedisTest{nc:c, p:p}
}

func testpanic(err error) {
	if err != nil {
		return
	}
	panic(err)
}


//string
func (r *RedisTest) StringT() {


	/*
	Get(key string) (string, error)
	Set(key, value string) error
	SetEx(key, value string, expire int) error
	SetNX(key, value string) (bool, error)
	Incr(key string) (int, error)

	//new
	Decr(key string) (int64, error)
	Append(key string, value string) (int64, error)
	IncrBy(key string, value int64) (int64, error)
	Strlen(key string) (int64, error)
	GetRange(key string, start int64, end int64) (string, error)
	MSet(fields map[string]interface{}) (string, error) //return ok
	MGet(fields ...string) ([]string, error)

	 */
	var values string
	//var valueint int
	var valueint64 int64
	//var valueb bool
	var err error

	fmt.Printf("Set--[rep:][err:%v]\n", r.nc.Set("redis.StringT.key", "redis.StringT.value"))
	values, err = r.nc.Get("redis.StringT.key")
	fmt.Printf("Get--[rep:%s][err:%v]\n", values, err)

	//Decr
	r.nc.Set("redis.Decr.key", "10")
	valueint64, err = r.nc.Decr("redis.Decr.key")
	fmt.Printf("Get--[rep:%d][err:%v] ex:9\n", valueint64, err)

	//append
	r.nc.Set("redis.Append.key", "redis.Append.key")
	valueint64, err = r.nc.Append("redis.Append.key", "append")
	values, err =  r.nc.Get("redis.Append.key")
	fmt.Printf("Append-- [req:%v][err:%v] ex:redis.Append.key.append\n", values, err)

	//Increby
	r.nc.Set("redis.IncreBy", "1")
	valueint64, err = r.nc.IncrBy("redis.IncreBy", 10)
	fmt.Printf("Increby--[req:%d][err:%v] ex:11\n", valueint64, err)

	//Strlen
	r.nc.Set("redis.Strlen", "Strlen")
	valueint64, err = r.nc.Strlen("redis.Strlen")
	fmt.Printf("Strlen--[%d][%v] ex:6\n", valueint64, err)

	//GetRange
	r.nc.Set("redis.GetRange", "redis.GetRange")
	values, err = r.nc.GetRange("redis.GetRange", 0, 5)
	fmt.Printf("GetRange--[%s][%v]\n", values, err)

	//Mset
	valumem := map[string]interface{}{
		"Mset01": "Mset01",
		"Mset02": "Mset01",
	}

	values, err = r.nc.MSet(valumem)
	fmt.Printf("Mset--[%v][%v]\n", values, err)
	var valueslices []string
	valueslices, err = r.nc.MGet("Mset01", "Mset02")
	fmt.Printf("MGET--[%v][%v] ex: [Mset01 Mset01]\n", valueslices, err)


}
//command
func (r *RedisTest) ComandT() {
/*
	//new
	Persist(key string) (bool, error)
	PExpire(key string, micro int64) (bool, error)
	PTTL(key string) (int64, error) // -2 -1  micro
	TYPE(key string) (string, error) //none string list set zset hash
	RandomKey() (string, error)
 */

 	var v64 int64
 	var b bool
 	var s string
 	var e error

 	r.nc.SetEx("SetEx", "SetEx", 100)
 	v64, e = r.nc.PTTL("SetEx")
 	fmt.Printf("PTTL--[%v][%v]\n", v64, e)

 	b, e = r.nc.Persist("SetEx")
 	fmt.Printf("Persist--[%v][%v]\n", b, e)

	b, e = r.nc.PExpire("SetEx", 100000)
	fmt.Printf("PExpire--[%v][%v]\n", b, e)

	s, e = r.nc.TYPE("SetEx")
	fmt.Printf("TYPE--[%v][%v]\n", s, e)

	s, e = r.nc.RandomKey()
	fmt.Printf("RandomKey--[%v][%v]\n", s, e)

	s, e = r.nc.Ping()
	fmt.Printf("PING--[%v][%v]", s, e)





 	}
//hash
func (r *RedisTest) HashT() {
/*
	//new
	HIncrByFloat(key string, field string, value float64) (newValue float64, err error)
	HSetNX(key string, field string, value string) (isSet bool, err error)
	HKeys(key string)([]string, error)
 */
	var b bool
	var e error
	var f64 float64

	b, e = r.nc.HSetNX("HSET", "float","11.11")
	fmt.Printf("HSetNX--[%v][%v] ex:true\n", b, e)

	f64, e = r.nc.HIncrByFloat("HSET", "float", 11.11)
	fmt.Printf("HIncrByFloat--[%v][%v] ex:22.22\n", f64, e)

	var slices []string
	r.nc.HSet("HSET", "int", "64")
	slices, e = r.nc.HKeys("HSET")
	fmt.Printf("HKeys--[%v][%v]\n ex: [float int]\n", slices, e)

}

// set
func (r *RedisTest) SetT() {
	/*
		//new
	SInter(keys ...string) ([]string, error)
	SInterStore(dest string, keys ...string) (int, error)
	SUnion(keys ...string) ([]string, error)
	SUnionStore(des string, keys ...string)(int, error)
	*/

	var slices []string
	var i int
	var e error

	r.nc.SAdd("SET1", "1","2","3","4")
	r.nc.SAdd("SET2", "4","5","6","7")

	slices, e = r.nc.SInter("SET1", "SET2")
	fmt.Printf("SInter--[%v][%v]\n", slices, e)

	i, e = r.nc.SInterStore("SET11", "SET1", "SET2")
	fmt.Printf("SInterStore--[%v][%v]\n", i, e)

	slices, e = r.nc.SUnion("SET1", "SET2")
	fmt.Printf("SUnion--[%v][%v]\n", slices, e)

	i, e = r.nc.SUnionStore("SET22", "SET1", "SET2")
	fmt.Printf("SUnionStore--[%v][%v]\n", i, e)


}

//管道测试
func (r *RedisTest) PipeLineT() {
	pipeline := func (t redis.Pipeline) {
		t.Set("k1", "v1")
		t.Set("k2", "v2")
		t.Set("k3", "v3")
		t.Get("k1")
		t.Get("k2")
		t.Get("k3")

	}
	value, err := r.nc.Pipelined(pipeline)
	fmt.Printf("Transation[%v][%v]", value, err)

	for _, v := range value {
		s, _ := redigo.String(v, nil)
		fmt.Printf("[%v]\n", s)
	}

}


//事务
func (r *RedisTest) TransationT() {
	transaction := func (t redis.Transaction) {
		t.Set("k1", "v1")
		t.Set("k2", "v2")
		t.Set("k3", "v3")
	}
	value, err := r.nc.Transaction(transaction)
	fmt.Printf("Transation[%v][%v]", value, err)
}



//pool
func (r *RedisTest) PoolT() {
	r.p.Set("H1", "H1")
	r.p.LPush("L", "init")
	r.p.RPush("R", "init")
	r.p.LPushX("L", "L1")
	r.p.RPushX("R", "R1")
	r.p.RPopLPush("L", "R")
	r.p.LSet("R", 1, "L11")


}

/*
func watch() {
	redigo.PubSubConn{}
}
*/

func (r *RedisTest) Scan() {

}

func (r *RedisTest) SScan() {
	key := "sscan"

	r.nc.SAdd(key, "a", "b", "c", "d", "e")

	var scanned []string
	var cursor int
	var matches []string
	var err error

	cursor, matches, err = r.nc.SScan(key, cursor, "", 1)
	scanned = append(scanned, matches...)
	for cursor != 0 {
		cursor, matches, err = r.nc.SScan(key, cursor, "", 1)
		scanned = append(scanned, matches...)
	}

	fmt.Printf("SSCAN--[%v][%v]\n", scanned, err)
}

func (r *RedisTest) ZScan() {
	key := "zscan"
	c := r.nc

	c.ZAdd(key, 1, "a")
	c.ZAdd(key, 2, "b")
	c.ZAdd(key, 3, "c")
	c.ZAdd(key, 4, "d")
	c.ZAdd(key, 5, "e")

	var scanned []string
	var scannedScores []float64
	var cursor int
	var matches []string
	var scores []float64
	var err error

	cursor, matches, scores, err = c.ZScan(key, cursor, "", 1)
	scanned = append(scanned, matches...)
	scannedScores = append(scannedScores, scores...)
	for cursor != 0 {
		cursor, matches, scores, err = c.ZScan(key, cursor, "", 1)
		scannedScores = append(scannedScores, scores...)
		scanned = append(scanned, matches...)
	}

	fmt.Printf("ZSCAN [%v][%v][%v]\n", scanned, scannedScores, err)
}

//list
func (r *RedisTest) ListT() {
/*
	//new
	RPopLPush(src string, dest string) (string, error)
	LSet(key string, index int, value string) error
	RPushX(key string, value string) (isPushed bool, err error)
	LPushX(key string, value string) (isPushed bool, err error)
 */

 	var s string
 	var b bool
 	var e error

 	r.nc.LPush("L", "init")
 	r.nc.RPush("R", "init")

 	b, e = r.nc.LPushX("L", "L1")
 	fmt.Printf("LPushX--[%v][%v] ex:true \n", b, e)

 	b, e = r.nc.RPushX("R", "R1")
	fmt.Printf("LPushX--[%v][%v] ex:true \n", b, e)

 	s, e = r.nc.RPopLPush("L", "R")
	fmt.Printf("RPopLPush--[%v][%v] ex:init \n", s, e)

 	e = r.nc.LSet("R", 1, "L11")
	fmt.Printf("LSet--[%v] \n", e)

}



func main() {
	// Using an arbitrary password should fallback to using no password
/*
	url := "redis://192.168.120.115:6379"
	parsedURL, _ := netURL.Parse(url)
	c, err := redis.NewConnection(parsedURL)
	if err != nil {
		panic(err)
	}
	err = c.Set("k1", "v1")
	if err != nil {
		fmt.Print("nil")
	}

	c.Close()
*/

	c := newConnectTest()
	c.StringT()
	c.ComandT()
	c.HashT()
	c.ListT()
	c.SetT()
	c.TransationT()
	c.PipeLineT()
	c.PoolT()
	c.SScan()
	c.ZScan()

	c.nc.Close()
	c.p.Shutdown()
}
