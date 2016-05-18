package ec2

import "github.com/timob/javabind"

type ServicesEc2ModelCreateRouteTableResultInterface interface {
	JavaLangObjectInterface

	// public void setRouteTable(com.amazonaws.services.ec2.model.RouteTable)
	SetRouteTable(a ServicesEc2ModelRouteTableInterface) 

	// public com.amazonaws.services.ec2.model.RouteTable getRouteTable()
	GetRouteTable() *ServicesEc2ModelRouteTable

	// public com.amazonaws.services.ec2.model.CreateRouteTableResult withRouteTable(com.amazonaws.services.ec2.model.RouteTable)
	WithRouteTable(a ServicesEc2ModelRouteTableInterface) *ServicesEc2ModelCreateRouteTableResult

	// public com.amazonaws.services.ec2.model.CreateRouteTableResult clone()
	Clone() *ServicesEc2ModelCreateRouteTableResult
}

type ServicesEc2ModelCreateRouteTableResult struct {
	JavaLangObject
}

// public com.amazonaws.services.ec2.model.CreateRouteTableResult()
func NewServicesEc2ModelCreateRouteTableResult() (*ServicesEc2ModelCreateRouteTableResult) {

	obj, err := javabind.GetEnv().NewObject("com/amazonaws/services/ec2/model/CreateRouteTableResult")
	if err != nil {
		panic(err)
	}
	x := &ServicesEc2ModelCreateRouteTableResult{}
	x.Callable = &javabind.Callable{obj}
	return x
}

// public void setRouteTable(com.amazonaws.services.ec2.model.RouteTable)
func (jbobject *ServicesEc2ModelCreateRouteTableResult) SetRouteTable(a ServicesEc2ModelRouteTableInterface)  {
	conv_a := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "setRouteTable", javabind.Void, conv_a.Value().Cast("com/amazonaws/services/ec2/model/RouteTable"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()

}

// public com.amazonaws.services.ec2.model.RouteTable getRouteTable()
func (jbobject *ServicesEc2ModelCreateRouteTableResult) GetRouteTable() *ServicesEc2ModelRouteTable {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getRouteTable", "com/amazonaws/services/ec2/model/RouteTable")
	if err != nil {
		panic(err)
	}
	retconv := javabind.NewJavaToGoCallable()
	dst := &javabind.Callable{}
	retconv.Dest(dst)
	if err := retconv.Convert(javabind.ObjectRef(jret)); err != nil {
		panic(err)
	}
	retconv.CleanUp()
	unique_x := &ServicesEc2ModelRouteTable{}
	unique_x.Callable = dst
	return unique_x
}

// public com.amazonaws.services.ec2.model.CreateRouteTableResult withRouteTable(com.amazonaws.services.ec2.model.RouteTable)
func (jbobject *ServicesEc2ModelCreateRouteTableResult) WithRouteTable(a ServicesEc2ModelRouteTableInterface) *ServicesEc2ModelCreateRouteTableResult {
	conv_a := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "withRouteTable", "com/amazonaws/services/ec2/model/CreateRouteTableResult", conv_a.Value().Cast("com/amazonaws/services/ec2/model/RouteTable"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()
	retconv := javabind.NewJavaToGoCallable()
	dst := &javabind.Callable{}
	retconv.Dest(dst)
	if err := retconv.Convert(javabind.ObjectRef(jret)); err != nil {
		panic(err)
	}
	retconv.CleanUp()
	unique_x := &ServicesEc2ModelCreateRouteTableResult{}
	unique_x.Callable = dst
	return unique_x
}

// public java.lang.String toString()
func (jbobject *ServicesEc2ModelCreateRouteTableResult) ToString() string {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "toString", "java/lang/String")
	if err != nil {
		panic(err)
	}
	retconv := javabind.NewJavaToGoString()
	dst := new(string)
	retconv.Dest(dst)
	if err := retconv.Convert(javabind.ObjectRef(jret)); err != nil {
		panic(err)
	}
	retconv.CleanUp()
	return *dst
}

// public boolean equals(java.lang.Object)
func (jbobject *ServicesEc2ModelCreateRouteTableResult) Equals(a interface{}) bool {
	conv_a := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "equals", javabind.Boolean, conv_a.Value().Cast("java/lang/Object"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()
	return jret.(bool)
}

// public int hashCode()
func (jbobject *ServicesEc2ModelCreateRouteTableResult) HashCode() int {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "hashCode", javabind.Int)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

// public com.amazonaws.services.ec2.model.CreateRouteTableResult clone()
func (jbobject *ServicesEc2ModelCreateRouteTableResult) Clone() *ServicesEc2ModelCreateRouteTableResult {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "clone", "com/amazonaws/services/ec2/model/CreateRouteTableResult")
	if err != nil {
		panic(err)
	}
	retconv := javabind.NewJavaToGoCallable()
	dst := &javabind.Callable{}
	retconv.Dest(dst)
	if err := retconv.Convert(javabind.ObjectRef(jret)); err != nil {
		panic(err)
	}
	retconv.CleanUp()
	unique_x := &ServicesEc2ModelCreateRouteTableResult{}
	unique_x.Callable = dst
	return unique_x
}

// public java.lang.Object clone() throws java.lang.CloneNotSupportedException
func (jbobject *ServicesEc2ModelCreateRouteTableResult) Clone2() (*JavaLangObject, error) {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "clone", "java/lang/Object")
	if err != nil {
		var zero *JavaLangObject
		return zero, err
	}
	retconv := javabind.NewJavaToGoCallable()
	dst := &javabind.Callable{}
	retconv.Dest(dst)
	if err := retconv.Convert(javabind.ObjectRef(jret)); err != nil {
		panic(err)
	}
	retconv.CleanUp()
	unique_x := &JavaLangObject{}
	unique_x.Callable = dst
	return unique_x, nil
}


