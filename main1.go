package main

import "fmt"

func main()  {
 persion1 := NewChinese()
 age := persion1.Age()
 if age < 30 {
 	fmt.Println("年龄上符合标准")
 }else {
 	fmt.Println("年龄上不符合")
 }
 fmt.Println(persion1)
}

type Person interface {
	Shanliang() bool
	WeiRenChuShi()
	Height() int
	Weight() int
	Age() int
	Salary() int
}

type Chinese struct {
	Name string
	Sex string
	IsShanliang bool
	High int
	Wei int
	AgeNum int
	Money int
	IsMajiang bool
}

func NewChinese() * Chinese{
	c := &Chinese{
		Name:        "龚江华",
		Sex:         "女",
		IsShanliang: true,
		High:        160,
		Wei:         180,
		AgeNum:      12,
		Money:       30,
		IsMajiang:   true,
	}
	return c
}

func (c *Chinese) Shanliang() bool  {
	return c.IsShanliang
}
func (c *Chinese) WeiRenChuShi() {
	fmt.Println(c.Name + "为人处世能力好")
}
func (c *Chinese) Height() int {
	return c.High
}
func (c *Chinese) Weight() int {
	return c.Wei
}
func (c *Chinese) Age() int {
	return c.AgeNum
}
func (c *Chinese) Salary() int {
	return c.Money
}
func (c *Chinese) Majiang() {
	if c.IsMajiang{
		fmt.Println(c.Name + "会打麻将")
	}else {
		fmt.Println(c.Name + "不会打麻将")
	}
}

type Japan struct {
	Name string
	AgeNum int
	IsShanLiang bool
	High int
	Weigh int
	Money int
	Go bool
}

func NewJapan() *Japan  {
	j := &Japan{
		Name:        "苍井空",
		AgeNum:      22,
		IsShanLiang: true,
		High:        163,
		Weigh:      90,
		Money:       100000,
		Go:          true,
	}
	return j
}
func (j *Japan) ShanLiang() bool {
	return j.IsShanLiang
}
func (j *Japan) WeiRenChuShi() {
	fmt.Println(j.Name + "为人处世能力好")
}
func (j *Japan) Height() int {
	return j.High
}
func (j *Japan) Weight() int {
	return j.Weigh
}
func (j *Japan) Age() int {
	return j.AgeNum
}
func (j *Japan) Salary() int {
	return j.Money
}
func (j *Japan) Eat(food string) {
	fmt.Println(j.Name + "喜欢吃:",food)
}
func (j *Japan) Gos() bool {
	return j.Go
}