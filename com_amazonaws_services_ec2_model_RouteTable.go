package ec2

import "github.com/timob/javabind"

type ServicesEc2ModelRouteTableInterface interface {
	JavaLangObjectInterface

	// public void setRouteTableId(java.lang.String)
	SetRouteTableId(a string) 

	// public java.lang.String getRouteTableId()
	GetRouteTableId() string

	// public com.amazonaws.services.ec2.model.RouteTable withRouteTableId(java.lang.String)
	WithRouteTableId(a string) *ServicesEc2ModelRouteTable

	// public void setVpcId(java.lang.String)
	SetVpcId(a string) 

	// public java.lang.String getVpcId()
	GetVpcId() string

	// public com.amazonaws.services.ec2.model.RouteTable withVpcId(java.lang.String)
	WithVpcId(a string) *ServicesEc2ModelRouteTable

	// public java.util.List<com.amazonaws.services.ec2.model.Route> getRoutes()
	GetRoutes() []*ServicesEc2ModelRoute

	// public void setRoutes(java.util.Collection<com.amazonaws.services.ec2.model.Route>)
	SetRoutes(a []*ServicesEc2ModelRoute) 

	// public com.amazonaws.services.ec2.model.RouteTable withRoutes(com.amazonaws.services.ec2.model.Route...)
	WithRoutes(a ...*ServicesEc2ModelRoute) *ServicesEc2ModelRouteTable

	// public com.amazonaws.services.ec2.model.RouteTable withRoutes(java.util.Collection<com.amazonaws.services.ec2.model.Route>)
	WithRoutes2(a []*ServicesEc2ModelRoute) *ServicesEc2ModelRouteTable

	// public java.util.List<com.amazonaws.services.ec2.model.RouteTableAssociation> getAssociations()
	GetAssociations() []*ServicesEc2ModelRouteTableAssociation

	// public void setAssociations(java.util.Collection<com.amazonaws.services.ec2.model.RouteTableAssociation>)
	SetAssociations(a []*ServicesEc2ModelRouteTableAssociation) 

	// public com.amazonaws.services.ec2.model.RouteTable withAssociations(com.amazonaws.services.ec2.model.RouteTableAssociation...)
	WithAssociations(a ...*ServicesEc2ModelRouteTableAssociation) *ServicesEc2ModelRouteTable

	// public com.amazonaws.services.ec2.model.RouteTable withAssociations(java.util.Collection<com.amazonaws.services.ec2.model.RouteTableAssociation>)
	WithAssociations2(a []*ServicesEc2ModelRouteTableAssociation) *ServicesEc2ModelRouteTable

	// public java.util.List<com.amazonaws.services.ec2.model.Tag> getTags()
	GetTags() []*ServicesEc2ModelTag

	// public void setTags(java.util.Collection<com.amazonaws.services.ec2.model.Tag>)
	SetTags(a []*ServicesEc2ModelTag) 

	// public com.amazonaws.services.ec2.model.RouteTable withTags(com.amazonaws.services.ec2.model.Tag...)
	WithTags(a ...*ServicesEc2ModelTag) *ServicesEc2ModelRouteTable

	// public com.amazonaws.services.ec2.model.RouteTable withTags(java.util.Collection<com.amazonaws.services.ec2.model.Tag>)
	WithTags2(a []*ServicesEc2ModelTag) *ServicesEc2ModelRouteTable

	// public java.util.List<com.amazonaws.services.ec2.model.PropagatingVgw> getPropagatingVgws()
	GetPropagatingVgws() []*ServicesEc2ModelPropagatingVgw

	// public void setPropagatingVgws(java.util.Collection<com.amazonaws.services.ec2.model.PropagatingVgw>)
	SetPropagatingVgws(a []*ServicesEc2ModelPropagatingVgw) 

	// public com.amazonaws.services.ec2.model.RouteTable withPropagatingVgws(com.amazonaws.services.ec2.model.PropagatingVgw...)
	WithPropagatingVgws(a ...*ServicesEc2ModelPropagatingVgw) *ServicesEc2ModelRouteTable

	// public com.amazonaws.services.ec2.model.RouteTable withPropagatingVgws(java.util.Collection<com.amazonaws.services.ec2.model.PropagatingVgw>)
	WithPropagatingVgws2(a []*ServicesEc2ModelPropagatingVgw) *ServicesEc2ModelRouteTable

	// public com.amazonaws.services.ec2.model.RouteTable clone()
	Clone() *ServicesEc2ModelRouteTable
}

type ServicesEc2ModelRouteTable struct {
	JavaLangObject
}

// public com.amazonaws.services.ec2.model.RouteTable()
func NewServicesEc2ModelRouteTable() (*ServicesEc2ModelRouteTable) {

	obj, err := javabind.GetEnv().NewObject("com/amazonaws/services/ec2/model/RouteTable")
	if err != nil {
		panic(err)
	}
	x := &ServicesEc2ModelRouteTable{}
	x.Callable = &javabind.Callable{obj}
	return x
}

// public void setRouteTableId(java.lang.String)
func (jbobject *ServicesEc2ModelRouteTable) SetRouteTableId(a string)  {
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
func (jbobject *ServicesEc2ModelRouteTable) GetRouteTableId() string {
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

// public com.amazonaws.services.ec2.model.RouteTable withRouteTableId(java.lang.String)
func (jbobject *ServicesEc2ModelRouteTable) WithRouteTableId(a string) *ServicesEc2ModelRouteTable {
	conv_a := javabind.NewGoToJavaString()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "withRouteTableId", "com/amazonaws/services/ec2/model/RouteTable", conv_a.Value().Cast("java/lang/String"))
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
	unique_x := &ServicesEc2ModelRouteTable{}
	unique_x.Callable = dst
	return unique_x
}

// public void setVpcId(java.lang.String)
func (jbobject *ServicesEc2ModelRouteTable) SetVpcId(a string)  {
	conv_a := javabind.NewGoToJavaString()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "setVpcId", javabind.Void, conv_a.Value().Cast("java/lang/String"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()

}

// public java.lang.String getVpcId()
func (jbobject *ServicesEc2ModelRouteTable) GetVpcId() string {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getVpcId", "java/lang/String")
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

// public com.amazonaws.services.ec2.model.RouteTable withVpcId(java.lang.String)
func (jbobject *ServicesEc2ModelRouteTable) WithVpcId(a string) *ServicesEc2ModelRouteTable {
	conv_a := javabind.NewGoToJavaString()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "withVpcId", "com/amazonaws/services/ec2/model/RouteTable", conv_a.Value().Cast("java/lang/String"))
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
	unique_x := &ServicesEc2ModelRouteTable{}
	unique_x.Callable = dst
	return unique_x
}

// public java.util.List<com.amazonaws.services.ec2.model.Route> getRoutes()
func (jbobject *ServicesEc2ModelRouteTable) GetRoutes() []*ServicesEc2ModelRoute {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getRoutes", "java/util/List")
	if err != nil {
		panic(err)
	}
	retconv := javabind.NewJavaToGoList(javabind.NewJavaToGoCallable())
	dst := new([]*ServicesEc2ModelRoute)
	retconv.Dest(dst)
	if err := retconv.Convert(javabind.ObjectRef(jret)); err != nil {
		panic(err)
	}
	retconv.CleanUp()
	return *dst
}

// public void setRoutes(java.util.Collection<com.amazonaws.services.ec2.model.Route>)
func (jbobject *ServicesEc2ModelRouteTable) SetRoutes(a []*ServicesEc2ModelRoute)  {
	conv_a := javabind.NewGoToJavaCollection(javabind.NewGoToJavaCallable())
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "setRoutes", javabind.Void, conv_a.Value().Cast("java/util/Collection"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()

}

// public com.amazonaws.services.ec2.model.RouteTable withRoutes(com.amazonaws.services.ec2.model.Route...)
func (jbobject *ServicesEc2ModelRouteTable) WithRoutes(a ...*ServicesEc2ModelRoute) *ServicesEc2ModelRouteTable {
	conv_a := javabind.NewGoToJavaObjectArray(javabind.NewGoToJavaCallable(), "com/amazonaws/services/ec2/model/Route")
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "withRoutes", "com/amazonaws/services/ec2/model/RouteTable", conv_a.Value().Cast("com/amazonaws/services/ec2/model/Route"))
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
	unique_x := &ServicesEc2ModelRouteTable{}
	unique_x.Callable = dst
	return unique_x
}

// public com.amazonaws.services.ec2.model.RouteTable withRoutes(java.util.Collection<com.amazonaws.services.ec2.model.Route>)
func (jbobject *ServicesEc2ModelRouteTable) WithRoutes2(a []*ServicesEc2ModelRoute) *ServicesEc2ModelRouteTable {
	conv_a := javabind.NewGoToJavaCollection(javabind.NewGoToJavaCallable())
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "withRoutes", "com/amazonaws/services/ec2/model/RouteTable", conv_a.Value().Cast("java/util/Collection"))
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
	unique_x := &ServicesEc2ModelRouteTable{}
	unique_x.Callable = dst
	return unique_x
}

// public java.util.List<com.amazonaws.services.ec2.model.RouteTableAssociation> getAssociations()
func (jbobject *ServicesEc2ModelRouteTable) GetAssociations() []*ServicesEc2ModelRouteTableAssociation {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getAssociations", "java/util/List")
	if err != nil {
		panic(err)
	}
	retconv := javabind.NewJavaToGoList(javabind.NewJavaToGoCallable())
	dst := new([]*ServicesEc2ModelRouteTableAssociation)
	retconv.Dest(dst)
	if err := retconv.Convert(javabind.ObjectRef(jret)); err != nil {
		panic(err)
	}
	retconv.CleanUp()
	return *dst
}

// public void setAssociations(java.util.Collection<com.amazonaws.services.ec2.model.RouteTableAssociation>)
func (jbobject *ServicesEc2ModelRouteTable) SetAssociations(a []*ServicesEc2ModelRouteTableAssociation)  {
	conv_a := javabind.NewGoToJavaCollection(javabind.NewGoToJavaCallable())
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "setAssociations", javabind.Void, conv_a.Value().Cast("java/util/Collection"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()

}

// public com.amazonaws.services.ec2.model.RouteTable withAssociations(com.amazonaws.services.ec2.model.RouteTableAssociation...)
func (jbobject *ServicesEc2ModelRouteTable) WithAssociations(a ...*ServicesEc2ModelRouteTableAssociation) *ServicesEc2ModelRouteTable {
	conv_a := javabind.NewGoToJavaObjectArray(javabind.NewGoToJavaCallable(), "com/amazonaws/services/ec2/model/RouteTableAssociation")
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "withAssociations", "com/amazonaws/services/ec2/model/RouteTable", conv_a.Value().Cast("com/amazonaws/services/ec2/model/RouteTableAssociation"))
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
	unique_x := &ServicesEc2ModelRouteTable{}
	unique_x.Callable = dst
	return unique_x
}

// public com.amazonaws.services.ec2.model.RouteTable withAssociations(java.util.Collection<com.amazonaws.services.ec2.model.RouteTableAssociation>)
func (jbobject *ServicesEc2ModelRouteTable) WithAssociations2(a []*ServicesEc2ModelRouteTableAssociation) *ServicesEc2ModelRouteTable {
	conv_a := javabind.NewGoToJavaCollection(javabind.NewGoToJavaCallable())
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "withAssociations", "com/amazonaws/services/ec2/model/RouteTable", conv_a.Value().Cast("java/util/Collection"))
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
	unique_x := &ServicesEc2ModelRouteTable{}
	unique_x.Callable = dst
	return unique_x
}

// public java.util.List<com.amazonaws.services.ec2.model.Tag> getTags()
func (jbobject *ServicesEc2ModelRouteTable) GetTags() []*ServicesEc2ModelTag {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getTags", "java/util/List")
	if err != nil {
		panic(err)
	}
	retconv := javabind.NewJavaToGoList(javabind.NewJavaToGoCallable())
	dst := new([]*ServicesEc2ModelTag)
	retconv.Dest(dst)
	if err := retconv.Convert(javabind.ObjectRef(jret)); err != nil {
		panic(err)
	}
	retconv.CleanUp()
	return *dst
}

// public void setTags(java.util.Collection<com.amazonaws.services.ec2.model.Tag>)
func (jbobject *ServicesEc2ModelRouteTable) SetTags(a []*ServicesEc2ModelTag)  {
	conv_a := javabind.NewGoToJavaCollection(javabind.NewGoToJavaCallable())
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "setTags", javabind.Void, conv_a.Value().Cast("java/util/Collection"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()

}

// public com.amazonaws.services.ec2.model.RouteTable withTags(com.amazonaws.services.ec2.model.Tag...)
func (jbobject *ServicesEc2ModelRouteTable) WithTags(a ...*ServicesEc2ModelTag) *ServicesEc2ModelRouteTable {
	conv_a := javabind.NewGoToJavaObjectArray(javabind.NewGoToJavaCallable(), "com/amazonaws/services/ec2/model/Tag")
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "withTags", "com/amazonaws/services/ec2/model/RouteTable", conv_a.Value().Cast("com/amazonaws/services/ec2/model/Tag"))
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
	unique_x := &ServicesEc2ModelRouteTable{}
	unique_x.Callable = dst
	return unique_x
}

// public com.amazonaws.services.ec2.model.RouteTable withTags(java.util.Collection<com.amazonaws.services.ec2.model.Tag>)
func (jbobject *ServicesEc2ModelRouteTable) WithTags2(a []*ServicesEc2ModelTag) *ServicesEc2ModelRouteTable {
	conv_a := javabind.NewGoToJavaCollection(javabind.NewGoToJavaCallable())
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "withTags", "com/amazonaws/services/ec2/model/RouteTable", conv_a.Value().Cast("java/util/Collection"))
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
	unique_x := &ServicesEc2ModelRouteTable{}
	unique_x.Callable = dst
	return unique_x
}

// public java.util.List<com.amazonaws.services.ec2.model.PropagatingVgw> getPropagatingVgws()
func (jbobject *ServicesEc2ModelRouteTable) GetPropagatingVgws() []*ServicesEc2ModelPropagatingVgw {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getPropagatingVgws", "java/util/List")
	if err != nil {
		panic(err)
	}
	retconv := javabind.NewJavaToGoList(javabind.NewJavaToGoCallable())
	dst := new([]*ServicesEc2ModelPropagatingVgw)
	retconv.Dest(dst)
	if err := retconv.Convert(javabind.ObjectRef(jret)); err != nil {
		panic(err)
	}
	retconv.CleanUp()
	return *dst
}

// public void setPropagatingVgws(java.util.Collection<com.amazonaws.services.ec2.model.PropagatingVgw>)
func (jbobject *ServicesEc2ModelRouteTable) SetPropagatingVgws(a []*ServicesEc2ModelPropagatingVgw)  {
	conv_a := javabind.NewGoToJavaCollection(javabind.NewGoToJavaCallable())
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "setPropagatingVgws", javabind.Void, conv_a.Value().Cast("java/util/Collection"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()

}

// public com.amazonaws.services.ec2.model.RouteTable withPropagatingVgws(com.amazonaws.services.ec2.model.PropagatingVgw...)
func (jbobject *ServicesEc2ModelRouteTable) WithPropagatingVgws(a ...*ServicesEc2ModelPropagatingVgw) *ServicesEc2ModelRouteTable {
	conv_a := javabind.NewGoToJavaObjectArray(javabind.NewGoToJavaCallable(), "com/amazonaws/services/ec2/model/PropagatingVgw")
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "withPropagatingVgws", "com/amazonaws/services/ec2/model/RouteTable", conv_a.Value().Cast("com/amazonaws/services/ec2/model/PropagatingVgw"))
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
	unique_x := &ServicesEc2ModelRouteTable{}
	unique_x.Callable = dst
	return unique_x
}

// public com.amazonaws.services.ec2.model.RouteTable withPropagatingVgws(java.util.Collection<com.amazonaws.services.ec2.model.PropagatingVgw>)
func (jbobject *ServicesEc2ModelRouteTable) WithPropagatingVgws2(a []*ServicesEc2ModelPropagatingVgw) *ServicesEc2ModelRouteTable {
	conv_a := javabind.NewGoToJavaCollection(javabind.NewGoToJavaCallable())
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "withPropagatingVgws", "com/amazonaws/services/ec2/model/RouteTable", conv_a.Value().Cast("java/util/Collection"))
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
	unique_x := &ServicesEc2ModelRouteTable{}
	unique_x.Callable = dst
	return unique_x
}

// public java.lang.String toString()
func (jbobject *ServicesEc2ModelRouteTable) ToString() string {
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
func (jbobject *ServicesEc2ModelRouteTable) Equals(a interface{}) bool {
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
func (jbobject *ServicesEc2ModelRouteTable) HashCode() int {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "hashCode", javabind.Int)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

// public com.amazonaws.services.ec2.model.RouteTable clone()
func (jbobject *ServicesEc2ModelRouteTable) Clone() *ServicesEc2ModelRouteTable {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "clone", "com/amazonaws/services/ec2/model/RouteTable")
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

// public java.lang.Object clone() throws java.lang.CloneNotSupportedException
func (jbobject *ServicesEc2ModelRouteTable) Clone2() (*JavaLangObject, error) {
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


