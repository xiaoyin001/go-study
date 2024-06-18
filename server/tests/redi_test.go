package tests

import (
	"context"
	"fmt"
	"github.com/redis/go-redis/v9"
	"strconv"
	"sync"
	"testing"
	"time"
)

// New一个 Redis 客户端的连接
func newRedis() *redis.Client {
	// 更多 Options 的参数解释可以参考 go Redis 的文档
	// https://redis.uptrace.dev/zh/guide/go-redis-option.html#redis-client

	return redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
	})
}

// Redis string类型操作例子
func TestRedisString(t *testing.T) {
	// 更多相关命令可以参考 菜鸟教程
	// https://www.runoob.com/redis/redis-strings.html

	mRedis := newRedis()
	mCtx := context.Background()

	// 通过 redis 命令的形式插入 string类型的KV
	mErr := mRedis.Do(mCtx, "Set", "Key_String_01", "Value_String_01").Err()
	if mErr != nil {
		fmt.Println("Key_String_01-", mErr.Error())
	}

	// 通过 redis 命令的形式插入 string类型的KV 并设置过期时间（单位：秒）
	mErr = mRedis.Do(mCtx, "SetEx", "Key_String_02", 30, "Value_String_02").Err()
	if mErr != nil {
		fmt.Println("Key_String_02-", mErr.Error())
	}

	// 通过 go-redis 提供的方法插入string类型数据并可以设置过去时间（这里设置过期时间30秒）
	mErr = mRedis.Set(mCtx, "Key_String_03", "Value_String_03", time.Second*30).Err()
	if mErr != nil {
		fmt.Println("Key_String_03-", mErr.Error())
	}

	// 通过 redis 命令形式获取指定Key的值 string类型的KV
	// 查询的Key不存在，个人建议在获取Key之前还是需要看一下Key是否存在
	mValue, mErr := mRedis.Do(mCtx, "Get", "Key_String_00").Result()
	if mErr != nil {
		fmt.Println("Key_String_00-查询错误-", mErr.Error())
	} else {
		fmt.Println("Key_String_00 =", mValue)
	}
	// 查询的Key存在
	mValue = mRedis.Do(mCtx, "Get", "Key_String_01").String()
	fmt.Println("Key_String_01 =", mValue)

	// 通过 go-redis 提供的方法 查询string类型数据
	mStrValue, mErr := mRedis.Get(mCtx, "Key_String_02").Result()
	if mErr != nil {
		fmt.Println("Key_String_02-查询错误-", mErr.Error())
	}
	fmt.Println("Key_String_02 =", mStrValue)
}

// Redis Hash类型操作例子
func TestRedisHash(t *testing.T) {
	// 更多命令可以参考 菜鸟教程
	// https://www.runoob.com/redis/redis-hashes.html

	mRedis := newRedis()
	mCtx := context.Background()

	// 命令的形式插入 Hash类型的数据
	mErr := mRedis.Do(mCtx, "HSet", "Key_Hash_01", "Value_Hash_01-01K", "Value_Hash_01_01V").Err()
	if mErr != nil {
		fmt.Println("Key_Hash_01-01K-插入错误-", mErr.Error())
	}
	mErr = mRedis.Do(mCtx, "HSet", "Key_Hash_01", "Value_Hash_01_02K", "Value_Hash_01_02V").Err()
	if mErr != nil {
		fmt.Println("Key_Hash_01-02K-插入错误-", mErr.Error())
	}
	// 这里覆盖上面插入 Value_Hash_02_K 的哈希数据
	mErr = mRedis.Do(mCtx, "HSet", "Key_Hash_01", "Value_Hash_01_02K", "Value_Hash_02_New").Err()
	if mErr != nil {
		fmt.Println("Key_Hash_01-02K-覆盖插入错误-", mErr.Error())
	}

	// go-redis 提供的方法插入Hash类型的数据
	mErr = mRedis.HSet(mCtx, "Key_Hash_02", "Key_Hash_02_01K", "Value_Hash_02_01V").Err()
	if mErr != nil {
		fmt.Println("Key_Hash_02_01K-插入错误-", mErr.Error())
	}

	// 设置过期时间
	mOk, mErr := mRedis.Expire(mCtx, "Key_Hash_01", time.Second*30).Result()
	if mErr != nil {
		fmt.Println("Key_Hash_01-设置过期时间错误-", mErr.Error())
	}
	if mOk {
		fmt.Println("Key_Hash_01-设置超时成功")
	} else {
		fmt.Println("Key_Hash_01-设置超时失败")
	}

	// 命令行的形式获取 Hash类型的数据
	mValue, mErr := mRedis.Do(mCtx, "HGet", "Key_Hash_01", "Value_Hash_01-01K").Result()
	if mErr != nil {
		fmt.Println("Value_Hash_01-01K-获取失败-", mErr.Error())
	} else {
		fmt.Println("Value_Hash_01-01K =", mValue)
	}

	// go-redis 提供的方法进行查询
	mValue, mErr = mRedis.HGet(mCtx, "Key_Hash_02", "Key_Hash_02_01K").Result()
	if mErr != nil {
		fmt.Println("Value_Hash_02-01K-获取失败-", mErr.Error())
	} else {
		fmt.Println("Value_Hash_02-01K =", mValue)
	}
}

var mIsSleep bool

// Redis List类型操作例子
func TestRedisList(t *testing.T) {
	// 更多命令可以参考 菜鸟教程
	// https://www.runoob.com/redis/redis-lists.html

	mRedis := newRedis()
	mCtx := context.Background()
	mIsSleep = false

	//  Lpush 命令将一个或多个值插入到列表头部 【LPUSH KEY_NAME VALUE1.. VALUEN】
	// 命令的形式插入 List类型数据
	mErr := mRedis.Do(mCtx, "LPush", "Key_List_01", "Value_List_01").Err()
	if mErr != nil {
		fmt.Println("Value_List_01-插入错误-", mErr.Error())
	}
	// 可以插入单个也可以插入多个数据
	mErr = mRedis.Do(mCtx, "LPush", "Key_List_01", "Value_List_02", "Value_List_03", "Value_List_04").Err()
	if mErr != nil {
		fmt.Println("Value_List_02~04-插入错误-", mErr.Error())
	}

	// 使用 go-redis 提供的方法插入 List类型数据（同样也可以插入多个数据）
	mErr = mRedis.LPush(mCtx, "Key_List_01", "Value_List_05").Err()
	if mErr != nil {
		fmt.Println("Value_List_05-插入错误-", mErr.Error())
	}

	// 设置超时时间
	mOk, mErr := mRedis.Expire(mCtx, "Key_List_01", time.Second*50).Result()
	if mErr != nil {
		fmt.Println("Key_List_01-设置超时错误-", mErr.Error())
	} else {
		if mOk {
			fmt.Println("Key_List_01-超时设置成功", mOk)
		} else {
			fmt.Println("Key_List_01-超时设置失败", mOk)
		}
	}

	if mIsSleep {
		time.Sleep(time.Second * 6)
	}

	// 使用命令的形式取出 List类型数据
	// Lpop 命令用于移除并返回列表的第一个元素
	mValue, mErr := mRedis.Do(mCtx, "LPop", "Key_List_01").Result()
	if mErr != nil {
		fmt.Println("LPop-Key_List_01-取出List中数据失败-", mErr.Error())
	} else {
		fmt.Println("LPop-Key_List_01 Value=", mValue)
	}

	if mIsSleep {
		time.Sleep(time.Second * 6)
	}

	// Rpop 命令用于移除列表的最后一个元素，返回值为移除的元素
	mValue, mErr = mRedis.Do(mCtx, "RPop", "Key_List_01").Result()
	if mErr != nil {
		fmt.Println("RPop-Key_List_01-取出List中数据失败-", mErr.Error())
	} else {
		fmt.Println("RPop-Key_List_01 Value=", mValue)
	}

	if mIsSleep {
		time.Sleep(time.Second * 6)
	}

	// 使用 go-redis 提供的方法进行插入 List类型数据
	mErr = mRedis.LPush(mCtx, "Key_List_01", "Value_List_06").Err()
	if mErr != nil {
		fmt.Println("Value_List_06-数据插入失败-", mErr.Error())
	}

	if mIsSleep {
		time.Sleep(time.Second * 6)
	}

	// 使用 go-redis 提供的方法取出 List类型数据
	// Rpop 命令用于移除列表的最后一个元素，返回值为移除的元素
	mValue, mErr = mRedis.RPop(mCtx, "Key_List_01").Result()
	if mErr != nil {
		fmt.Println("RPop()-Key_List_01-取出List中数据失败-", mErr.Error())
	} else {
		fmt.Println("RPop()-Key_List_01 Value=", mValue)
	}

	if mIsSleep {
		time.Sleep(time.Second * 6)
	}

	// Lpop 命令用于移除并返回列表的第一个元素
	mValue, mErr = mRedis.LPop(mCtx, "Key_List_01").Result()
	if mErr != nil {
		fmt.Println("LPop()-Key_List_01-取出List中数据失败-", mErr.Error())
	} else {
		fmt.Println("LPop()-Key_List_01 Value=", mValue)
	}

	/*
		总结：
			1、Push数据不存在覆盖这一说，就如同Push的意思一样，往指定Key下面的List中追加数据
			2、在Push数据的时候看是需要将数据从头插入还是插入在尾部，L对应前、R对应尾
			3、在Pop取数据的时候也同上类似，也有L、R的区分，代表这从头还是尾取出数据
	*/
}

// Redis Set（集合）类型操作例子
func TestRedisSet(t *testing.T) {
	// 更多命令可以参考 菜鸟教程
	// https://www.runoob.com/redis/redis-sets.html

	mRedis := newRedis()
	mCtx := context.Background()

	// 通过命令的形式插入 集合类型数据
	// SAdd 命令将一个或多个成员元素加入到集合中，已经存在于集合的成员元素将被忽略
	mErr := mRedis.Do(mCtx, "SAdd", "Key_Set_01", "Value_Set_01", "Value_Set_02").Err()
	if mErr != nil {
		fmt.Println("Value_Set_01~02-Add失败-", mErr.Error())
	}

	// 通过 go-redis 提供的方法进行Add数据
	mErr = mRedis.SAdd(mCtx, "Key_Set_01", "Value_Set_03", "Value_Set_04").Err()
	if mErr != nil {
		fmt.Println("Value_Set_03~04-Add失败-", mErr.Error())
	}

	for i := 0; i < 10; i++ {
		mValue := "Value_Set_0" + strconv.Itoa(i)
		mRedis.SAdd(mCtx, "Key_Set_01", mValue)
	}

	// 通过命令的形式获取 集合类型的数据
	mValue, mErr := mRedis.Do(mCtx, "SMembers", "Key_Set_01").Result()
	if mErr != nil {
		fmt.Println("SMembers-Key_Set_01-查询集合数据错误-", mErr.Error())
	} else {
		fmt.Println("SMembers-Key_Set_01 Value=", mValue)
	}

	// 通过 go-redis 提供的方法进行查询 集合类型数据
	mValue, mErr = mRedis.SMembers(mCtx, "Key_Set_01").Result()
	if mErr != nil {
		fmt.Println("SMembers()-Key_Set_01-查询集合数据错误-", mErr.Error())
	} else {
		fmt.Println("SMembers()-Key_Set_01 Value=", mValue)
	}

	// 通过命令设置过期时间
	// 具体的可以参考菜鸟教程：https://www.runoob.com/redis/keys-expire.html
	mRes, mErr := mRedis.Do(mCtx, "Expire", "Key_Set_01", 30).Int64()
	if mErr != nil {
		fmt.Println("Key_Set_01-设置超时错误-", mErr.Error())
	} else {
		if mRes == 1 {
			fmt.Println("Key_Set_01-设置超时成功")
		} else {
			fmt.Println("Key_Set_01-设置超时失败")
		}
	}

	/*
		总结：添加的数据并不是按照插入的顺序进行存放的，存放的数据是无序且不重复
	*/
}

// Redis ZSte（有序集合）类型操作例子
func TestRedisZSet(t *testing.T) {
	// 更多命令可以参考 菜鸟教程
	// https://www.runoob.com/redis/redis-sorted-sets.html

	mRedis := newRedis()
	mCtx := context.Background()

	mItem := redis.Z{}
	// 通过命令的形式添加 有序集合类型数据
	// ZAdd 命令用于将一个或多个成员元素及其分数值加入到有序集当中
	// 已经是有序集合的成员，那么更新这个成员的分数值，并通过重新插入这个成员元素，来保证该成员在正确的位置上
	mErr := mRedis.Do(mCtx, "ZAdd", "Key_ZSet_01", 10, "Value_ZSet_00").Err()
	if mErr != nil {
		fmt.Println("Value_ZSet_00:10-添加失败-", mErr.Error())
	}

	// 通过 go-redis 提供的方法添加 有序集合类型数据
	mItem.Score = 10
	mItem.Member = "Value_ZSet_01"
	mErr = mRedis.ZAdd(mCtx, "Key_ZSet_01", mItem).Err()
	if mErr != nil {
		fmt.Println("Key_ZSet_01-添加数据错误-", mErr.Error())
	}

	mItem.Score = 9
	mItem.Member = "Value_ZSet_02"
	mRedis.ZAdd(mCtx, "Key_ZSet_01", mItem)
	mItem.Member = "Value_ZSet_03"
	mRedis.ZAdd(mCtx, "Key_ZSet_01", mItem)
	mItem.Member = "Value_ZSet_04"
	mRedis.ZAdd(mCtx, "Key_ZSet_01", mItem)

	mItem.Score = 11
	mItem.Member = "Value_ZSet_01"
	mRedis.ZAdd(mCtx, "Key_ZSet_01", mItem)

	// 命令方式获取 有序集合指定区间（排序下标）内的成员
	// ZRevRange 返回有序集中，指定区间内的成员，成员的位置按分数值递减(从大到小)来排列
	mValue, mErr := mRedis.Do(mCtx, "ZRevRange", "Key_ZSet_01", 0, 2).Result()
	if mErr != nil {
		fmt.Println("ZRevRange-Key_ZSet_01-获取指定区间内成员数据错误-", mErr.Error())
	} else {
		fmt.Println("ZRevRange-Key_ZSet_01-获取指定区间内成员数据- Value=", mValue)
	}

	// 通过 go-redis 提供的方法进行
	// ZRevRange 返回有序集中，指定区间内的成员，成员的位置按分数值递减(从大到小)来排列
	mValue, mErr = mRedis.ZRevRange(mCtx, "Key_ZSet_01", 0, 2).Result()
	if mErr != nil {
		fmt.Println("ZRevRange()-Key_ZSet_01-获取指定区间内成员数据错误-", mErr.Error())
	} else {
		fmt.Println("ZRevRange()-Key_ZSet_01-获取指定区间内成员数据- Value=", mValue)
	}

	/*
		总结：
			1、插入的数据排序是从大到小
			2、插入的数据跟集合一样，里面的成员是不可重复的，但是这里的 Score（可以理解为权重分数，越大排名越靠前） 是可以重复的
	*/
}

// 测试多线程同时操作Redis相同Key的数据
func TestDemo01(t *testing.T) {
	mRedis := newRedis()
	mCtx := context.Background()

	mWG := sync.WaitGroup{}
	mWG.Add(2)

	go func() {
		defer mWG.Done()

		for i := 0; i < 10; i++ {
			mRes, mErr := mRedis.Set(mCtx, "TestKey", "TestValue_"+strconv.Itoa(i), 0).Result()
			if mErr != nil {
				fmt.Println("-------------插入数据错误 i=", i, " Err =", mErr.Error())
			} else {
				fmt.Println("+++++++++++++插入数据成功 i=", i, " Res =", mRes)
			}

			//time.Sleep(time.Second)
		}
	}()

	go func() {
		defer mWG.Done()

		for i := 0; i < 10; i++ {
			mValue, mErr := mRedis.Get(mCtx, "TestKey").Result()
			if mErr != nil {
				fmt.Println("////////////读取数据错误 i=", i, " Err =", mErr.Error())
			} else {
				fmt.Println("============读取数据成功 i=", i, " Value=", mValue)
			}

			//time.Sleep(time.Second)
		}
	}()

	mWG.Wait()
}
