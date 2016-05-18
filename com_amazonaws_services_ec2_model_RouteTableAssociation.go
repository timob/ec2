package ec2

import "github.com/timob/javabind"

type ServicesEc2ModelRouteTableAssociationInterface interface {
	JavaLangObjectInterface

	// public void setRouteTableAssociationId(java.lang.String)
	SetRouteTableAssociationId(a string) 

	// public java.lang.String getRouteTableAssociationId()
	GetRouteTableAssociationId() string

	// public com.amazonaws.services.ec2.model.RouteTableAssociation withRouteTableAssociationId(java.lang.String)
	WithRouteTableAssociationId(a string) *ServicesEc2ModelRouteTableAssociation

	// public void setRouteTableId(java.lang.String)
	SetRouteTableId(a string) 

	// public java.lang.String getRouteTableId()
	GetRouteTableId() string

	// public com.amazonaws.services.ec2.model.RouteTableAssociation withRouteTableId(java.lang.String)
	WithRouteTableId(a string) *ServicesEc2ModelRouteTableAssociation

	// public void setSubnetId(java.lang.String)
	SetSubnetId(a string) 

	// public java.lang.String getSubnetId()
	GetSubnetId() string

	// public com.amazonaws.services.ec2.model.RouteTableAssociation withSubnetId(java.lang.String)
	WithSubnetId(a string) *ServicesEc2ModelRouteTableAssociation

	// public void setMain(java.lang.Boolean)
	SetMain(a bool) 

	// public java.lang.Boolean getMain()
	GetMain() bool

	// public com.amazonaws.services.ec2.model.RouteTableAssociation withMain(java.lang.Boolean)
	WithMain(a bool) *ServicesEc2ModelRouteTableAssociation

	// public java.lang.Boolean isMain()
	IsMain() bool

	// public com.amazonaws.services.ec2.model.RouteTableAssociation clone()
	Clone() *ServicesEc2ModelRouteTableAssociation
}

type ServicesEc2ModelRouteTableAssociation struct {
	JavaLangObject
}

// public com.amazonaws.services.ec2.model.RouteTableAssociation()
func NewServicesEc2ModelRouteTableAssociation() (*ServicesEc2ModelRouteTableAssociation) {

	obj, err := javabind.GetEnv().NewObject("com/amazonaws/services/ec2/model/RouteTableAssociation")
	if err != nil {
		panic(err)
	}
	x := &ServicesEc2ModelRouteTableAssociation{}
	x.Callable = &javabind.Callable{obj}
	return x
}

// public void setRouteTableAssociationId(java.lang.String)
func (jbobject *ServicesEc2ModelRouteTableAssociation) SetRouteTableAssociationId(a string)  {
	conv_a := javabind.NewGoToJavaString()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "setRouteTableAssociationId", javabind.Void, conv_a.Value().Cast("java/lang/String"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()

}

// public java.lang.String getRouteTableAssociationId()
func (jbobject *ServicesEc2ModelRouteTableAssociation) GetRouteTableAssociationId() string {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getRouteTableAssociationId", "java/lang/String")
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

// public com.amazonaws.services.ec2.model.RouteTableAssociation withRouteTableAssociationId(java.lang.String)
func (jbobject *ServicesEc2ModelRouteTableAssociation) WithRouteTableAssociationId(a string) *ServicesEc2ModelRouteTableAssociation {
	conv_a := javabind.NewGoToJavaString()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "withRouteTableAssociationId", "com/amazonaws/services/ec2/model/RouteTableAssociation", conv_a.Value().Cast("java/lang/String"))
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
	unique_x := &ServicesEc2ModelRouteTableAssociation{}
	unique_x.Callable = dst
	return unique_x
}

// public void setRouteTableId(java.lang.String)
func (jbobject *ServicesEc2ModelRouteTableAssociation) SetRouteTableId(a string)  {
	conv_a := javabind.NewGoToJavaString()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "setRouteTableId", javabind.Void, conv_a.Value().Cast("java/lang/String"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()

}

// public java.lang.String getRouteTableId()
func (jbobject *ServicesEc2ModelRouteTableAssociation) GetRouteTableId() string {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getRouteTableId", "java/lang/String")
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

// public com.amazonaws.services.ec2.model.RouteTableAssociation withRouteTableId(java.lang.String)
func (jbobject *ServicesEc2ModelRouteTableAssociation) WithRouteTableId(a string) *ServicesEc2ModelRouteTableAssociation {
	conv_a := javabind.NewGoToJavaString()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "withRouteTableId", "com/amazonaws/services/ec2/model/RouteTableAssociation", conv_a.Value().Cast("java/lang/String"))
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
	unique_x := &ServicesEc2ModelRouteTableAssociation{}
	unique_x.Callable = dst
	return unique_x
}

// public void setSubnetId(java.lang.String)
func (jbobject *ServicesEc2ModelRouteTableAssociation) SetSubnetId(a string)  {
	conv_a := javabind.NewGoToJavaString()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "setSubnetId", javabind.Void, conv_a.Value().Cast("java/lang/String"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()

}

// public java.lang.String getSubnetId()
func (jbobject *ServicesEc2ModelRouteTableAssociation) GetSubnetId() string {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getSubnetId", "java/lang/String")
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

// public com.amazonaws.services.ec2.model.RouteTableAssociation withSubnetId(java.lang.String)
func (jbobject *ServicesEc2ModelRouteTableAssociation) WithSubnetId(a string) *ServicesEc2ModelRouteTableAssociation {
	conv_a := javabind.NewGoToJavaString()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "withSubnetId", "com/amazonaws/services/ec2/model/RouteTableAssociation", conv_a.Value().Cast("java/lang/String"))
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
	unique_x := &ServicesEc2ModelRouteTableAssociation{}
	unique_x.Callable = dst
	return unique_x
}

// public void setMain(java.lang.Boolean)
func (jbobject *ServicesEc2ModelRouteTableAssociation) SetMain(a bool)  {
	conv_a := javabind.NewGoToJavaBoolean()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "setMain", javabind.Void, conv_a.Value().Cast("java/lang/Boolean"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()

}

// public java.lang.Boolean getMain()
func (jbobject *ServicesEc2ModelRouteTableAssociation) GetMain() bool {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getMain", "java/lang/Boolean")
	if err != nil {
		panic(err)
	}
	retconv := javabind.NewJavaToGoBoolean()
	dst := new(bool)
	retconv.Dest(dst)
	if err := retconv.Convert(javabind.ObjectRef(jret)); err != nil {
		panic(err)
	}
	retconv.CleanUp()
	return *dst
}

// public com.amazonaws.services.ec2.model.RouteTableAssociation withMain(java.lang.Boolean)
func (jbobject *ServicesEc2ModelRouteTableAssociation) WithMain(a bool) *ServicesEc2ModelRouteTableAssociation {
	conv_a := javabind.NewGoToJavaBoolean()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "withMain", "com/amazonaws/services/ec2/model/RouteTableAssociation", conv_a.Value().Cast("java/lang/Boolean"))
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
	unique_x := &ServicesEc2ModelRouteTableAssociation{}
	unique_x.Callable = dst
	return unique_x
}

// public java.lang.Boolean isMain()
func (jbobject *ServicesEc2ModelRouteTableAssociation) IsMain() bool {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "isMain", "java/lang/Boolean")
	if err != nil {
		panic(err)
	}
	retconv := javabind.NewJavaToGoBoolean()
	dst := new(bool)
	retconv.Dest(dst)
	if err := retconv.Convert(javabind.ObjectRef(jret)); err != nil {
		panic(err)
	}
	retconv.CleanUp()
	return *dst
}

// public java.lang.String toString()
func (jbobject *ServicesEc2ModelRouteTableAssociation) ToString() string {
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
func (jbobject *ServicesEc2ModelRouteTableAssociation) Equals(a interface{}) bool {
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
func (jbobject *ServicesEc2ModelRouteTableAssociation) HashCode() int {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "hashCode", javabind.Int)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

// public com.amazonaws.services.ec2.model.RouteTableAssociation clone()
func (jbobject *ServicesEc2ModelRouteTableAssociation) Clone() *ServicesEc2ModelRouteTableAssociation {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "clone", "com/amazonaws/services/ec2/model/RouteTableAssociation")
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
	unique_x := &ServicesEc2ModelRouteTableAssociation{}
	unique_x.Callable = dst
	return unique_x
}

// public java.lang.Object clone() throws java.lang.CloneNotSupportedException
func (jbobject *ServicesEc2ModelRouteTableAssociation) Clone2() (*JavaLangObject, error) {
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


