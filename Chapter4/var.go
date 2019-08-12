package chapter4

import (
	"fmt"
	"os"
	"runtime"
)

//声明变量的一般形式是使用var 关键字
//var identifier type
//不同于const不能省略type，注意不能省略这只是针对只声明变量，而const声明的同时必须赋值
//同时不同于常量，可以先只声明而不初始化
//const dddd int //error

//与许多语言不同，它在声明变量时将变量的类型放在变量的名称之后，
//这是为避免C那样含糊不清的声明方式，int* a,b 在这个 声明中
//只有a是指针，而b是一个int
//而在Go中可以都声明为指针类型
//var a,b *int
//也可使用因式分解的写法
//注意：一个包中不能有重名的全局变量，即使是一个包中的不同文件也不允许出现同名
var(
	aa int
	b bool
)
func notForError(){
	fmt.Print(aa,b)
}
//注意：不同于C++，与Java相同，当一个变量被声明后，系统自动赋予它该类型的零值
//int = 0 float = 0.0 bool = false
//string 为 空字符串
//指针为 nil
//记住，所有的内存在Go中都是经过初始化的
//变量的命名最好骆驼命名法

//一个变量在程序中都有一定性的作用范围,叫做作用域。如果一个变量在函数体外声明。则
//被认为是全局变量，可以在整个包甚至外部包（前提是被导出）使用，不管你声明在哪个源文件里或在哪个源
//文件里调用该变量
//在函数体内声明的变量称为局部变量，它们的作用域只在函数体内，参数和返回值变量也是局部变量
//一般情况下局部变量的作用域可以通过代码块（大括号括起来的部分）判断

//尽管变量的标识符必须是唯一的，但可以在某个代码块的内层代码块中使用相同名称的变量
//则此时外部的同名变量将会被暂时隐藏（结束内部代码后隐藏的外部同名变量又会出现，而内部同名变量则被释放）
//任何的操作都只会影响内部代码块的局部变量
//
//变量可以编译期间就被赋值，赋值给变量使用运算符等号 = ，也可以在运行时对变量进行赋值操作

//a = 15
//b = false

//当然可以将声明与赋值写在一起
//var identifier [type] = value //与常量const的声明一样
//Go的编译器可以根据变量的值来自动推断其类型，这与ruby与python这类动态语言类似
//但不同的是在编译时就已经完成了推断过程

//var a = 15
//var b = false
//var str = "dd"

//或使用因式分解方式
//var{
// a = 15
// b = 4
// }

//不过自动推断类型并不是任何时候都适用的，当想要给变量的类型并不是自动推断出的某种类型时
//仍需要显式给出
//var n int64 = 2
//像var a 这种语法是不对的
//变量的类型也可以在运行时实现自动推断
//var{
// HOME = os.Getenv("HOME")
// USER = os.Getenv("USER")
// }

//注意这种写法主要用于声明包级别的全局变量
//当在函数体内声明局部变量时，应使用简短声明的语法:=
//a:=1

func TestGoos(){
	var goos = runtime.GOOS
	fmt.Printf("The opearting system is: %s\n",goos)
	path := os.Getenv("PATH")
	fmt.Println("Path is "+path)
}


//值的类型和引用类型
//程序中所用到的内存在计算机中使用一堆箱子来表示，即“字”
//根据不同的处理器以及操作系统类型，所有的字都有32位（4字节）或64位（8字节）的相同长度
//所有的字都使用相关的内存地址来进行表示（以16进制数表示）

//所有像int float bool string这些基本类型都属于值类型
//使用这些类型的变量直接指向存在内存中的值
//另外，像数组和结构这些复合类型也是值类型
//当使用等号=将一个变量的值赋值给另一个变量时，如 j = i 实际上是在内存中将i值进行了拷贝
//可以通过使用&i来获取变量i的内存地址，例如：0xf84000000
//值类型的变量的值存储在栈中（注意：因为不是new的编译时就能确定有这么一个变量）
//
//内存地址会根据机器的不同而有所不同，甚至相同的程序，每台机器可能有不同的存储器布局
//更复杂的数据通常会需要使用多个字，这些数据一般使用引用类型保存

//一个引用类型的变量存储的是值所在的内存地址或内存地址中第一个字所在的位置
//这个内存地址被称为指针
//同一个引用类型的指针指向的多个字可以是在连续的内存地址中（内存布局是连续的），效率也最高
//也可以将这些字分散存放在内存中，每个字都指示了下个字所在的内存地址
//
//在Go语言中，指针属于引用类型，其它的引用类型还包括slices maps channel 被引用的变量会存储在堆中
//以便进行垃圾回收，且比栈拥有更大的内存空间
//
//打印
//函数Printf可以在fmt包外部使用，因为它以大写字母P开关
//用于打印输出到控制台

//这个格式化字符串可以含有一个或多个的格式化标识符 %...
//%s 代表字符串 %v代表使用类型的默认格式的标识符
//这些标识符所对应的值从格式化字符串后的第一个逗号开始按相同顺序添加

//函数Sprintf与Printf作用完全相同，只是前者将格式化后的字符串以返回值的方法返回给调用者
//函数fmt.Print和fmt.Println会自动使用格式化标识符%v对字符串进行格式化
//两者都会在每个参数之间自动增加空格，而后者最后还会加上一个换行符
func TestPint(){
	a:="abc:"
	b:=123
	fmt.Printf("a=%s,b=%d",a,b)
	fmt.Println(a,b)//有一个空格
	fmt.Println("abc:",123)//有一个空格
	//fmt的Println绝对遵守在两个参数之间添加空格的要求
	fmt.Println(123,"abc: ","=",123,4,"-",5)//有两个空格
	//然而Print不会这样，Print方法会在连续的数字型参数之间插入空格(目前观察到的)
	fmt.Print(123,"abc: ","=",123,4,"-",5)
}

//简短形式，使用:=赋值操作符
//因为变量的初始化时省略变量的类型而由系统自动推断，而这个时候再加上var有些多余
//因此简写为a:=50
//其类型由编译器自动推断
//这是使用变量的首选形式，但是它只能被用在函数体内，而不可用于全局变量的声明与赋值
//使用:=可以高效地创建一个新的变量，称为初始化声明
