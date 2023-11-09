package main

import (
	"fmt"
	"testing"
)

func TestMonster_Store(t *testing.T) {
	monster := &Monster{
		Name:  "牛魔王",
		Age:   12,
		Skill: "吐火",
	}

	res := monster.Store()
	if !res {
		t.Fatalf("monster.Store() 错误, 希望为=%v, 实际为=%v", true, res)
	}
	t.Logf("测试成功")
}

func TestMonster_ReStore(t *testing.T) {
	var monster Monster
	res := monster.ReStore()
	fmt.Println(res)
	if !res {
		t.Fatalf("monster.ReStore() 错误, 希望为=%v, 实际为=%v", true, res)
	}
	if monster.Name != "牛魔王" {
		t.Fatalf("monster.ReStore() 错误, 希望为=%v, 实际为=%v", "牛魔王", monster.Name)
	}
	t.Logf("测试成功")
}
