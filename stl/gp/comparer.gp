//#GOGP_IGNORE_BEGIN
///////////////////////////////////////////////////////////////////
//
// !!!!!!!!!!!! NEVER MODIFY THIS FILE MANUALLY !!!!!!!!!!!!
//
// This file was auto-generated by tool [github.com/vipally/gogp]
// Last update at: [Tue Oct 25 2016 16:15:22]
// Generate from:
//   [github.com/vipally/gx/stl/gp/comparer.go]
//   [github.com/vipally/gx/stl/gp/gp.gpg] [GOGP_REVERSE_comparer]
//
// Tool [github.com/vipally/gogp] info:
// CopyRight 2016 @Ally Dale. All rights reserved.
// Author  : Ally Dale(vipally@gmail.com)
// Blog    : http://blog.csdn.net/vipally
// Site    : https://github.com/vipally
// BuildAt : [Oct  8 2016 10:34:35]
// Version : 3.0.0.final
// 
///////////////////////////////////////////////////////////////////
//#GOGP_IGNORE_END





//

type Comparer<VALUE_TYPE> interface {
	F(left, right <VALUE_TYPE>) bool
}

type Comparer<VALUE_TYPE>Creator int

const (
	LESSER_<VALUE_TYPE> Comparer<VALUE_TYPE>Creator = iota
	GREATER_<VALUE_TYPE>
)

func (me Comparer<VALUE_TYPE>Creator) Create() (cmp Comparer<VALUE_TYPE>) {
	switch me {
	case LESSER_<VALUE_TYPE>:
		cmp = Lesser<VALUE_TYPE>(0)
	case GREATER_<VALUE_TYPE>:
		cmp = Greater<VALUE_TYPE>(0)
	}
	return
}

type Lesser<VALUE_TYPE> byte

func (this Lesser<VALUE_TYPE>) F(left, right <VALUE_TYPE>) (ok bool) {
	//#GOGP_IFDEF GOGP_HasLess
	ok = left.Less(right)
	//#GOGP_ELSE
	ok = left < right
	//#GOGP_ENDIF
	return
}

type Greater<VALUE_TYPE> byte

func (this Greater<VALUE_TYPE>) F(left, right <VALUE_TYPE>) (ok bool) {
	//#GOGP_IFDEF GOGP_HasGreat
	ok = left.Great(right)
	//#GOGP_ELSE
	ok = left > right
	//#GOGP_ENDIF
	return
}
