package ec2

import "github.com/timob/javabind"

type ServicesEc2ModelPrefixListInterface interface {
	JavaLangObjectInterface

	// public void setPrefixListId(java.lang.String)
	SetPrefixListId(a string) 

	// public java.lang.String getPrefixListId()
	GetPrefixListId() string

	// public com.amazonaws.services.ec2.model.PrefixList withPrefixListId(java.lang.String)
	WithPrefixListId(a string) *ServicesEc2ModelPrefixList

	// public void setPrefixListName(java.lang.String)
	SetPrefixListName(a string) 

	// public java.lang.String getPrefixListName()
	GetPrefixListName() string

	// public com.amazonaws.services.ec2.model.PrefixList withPrefixListName(java.lang.String)
	WithPrefixListName(a string) *ServicesEc2ModelPrefixList

	// public java.util.List<java.lang.String> getCidrs()
	GetCidrs() []string

	// public void setCidrs(java.util.Collection<java.lang.String>)
	SetCidrs(a []string) 

	// public com.amazonaws.services.ec2.model.PrefixList withCidrs(java.lang.String...)
	WithCidrs(a ...string) *ServicesEc2ModelPrefixList

	// public com.amazonaws.services.ec2.model.PrefixList withCidrs(java.util.Collection<java.lang.String>)
	WithCidrs2(a []string) *ServicesEc2ModelPrefixList

	// public com.amazonaws.services.ec2.model.PrefixList clone()
	Clone() *ServicesEc2ModelPrefixList
}

type ServicesEc2ModelPrefixList struct {
	JavaLangObject
}

// public com.amazonaws.services.ec2.model.PrefixList()
func NewServicesEc2ModelPrefixList() (*ServicesEc2ModelPrefixList) {

	obj, err := javabind.GetEnv().NewObject("com/amazonaws/services/ec2/model/PrefixList")
	if err != nil {
		panic(err)
	}
	x := &ServicesEc2ModelPrefixList{}
	x.Callable = &javabind.Callable{obj}
	return x
}

// public void setPrefixListId(java.lang.String)
func (jbobject *ServicesEc2ModelPrefixList) SetPrefixListId(a string)  {
	conv_a := javabind.NewGoToJavaString()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "setPrefixListId", javabind.Void, conv_a.Value().Cast("java/lang/String"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()

}

// public java.lang.String getPrefixListId()
func (jbobject *ServicesEc2ModelPrefixList) GetPrefixListId() string {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getPrefixListId", "java/lang/String")
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

// public com.amazonaws.services.ec2.model.PrefixList withPrefixListId(java.lang.String)
func (jbobject *ServicesEc2ModelPrefixList) WithPrefixListId(a string) *ServicesEc2ModelPrefixList {
	conv_a := javabind.NewGoToJavaString()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "withPrefixListId", "com/amazonaws/services/ec2/model/PrefixList", conv_a.Value().Cast("java/lang/String"))
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
	unique_x := &ServicesEc2ModelPrefixList{}
	unique_x.Callable = dst
	return unique_x
}

// public void setPrefixListName(java.lang.String)
func (jbobject *ServicesEc2ModelPrefixList) SetPrefixListName(a string)  {
	conv_a := javabind.NewGoToJavaString()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "setPrefixListName", javabind.Void, conv_a.Value().Cast("java/lang/String"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()

}

// public java.lang.String getPrefixListName()
func (jbobject *ServicesEc2ModelPrefixList) GetPrefixListName() string {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getPrefixListName", "java/lang/String")
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

// public com.amazonaws.services.ec2.model.PrefixList withPrefixListName(java.lang.String)
func (jbobject *ServicesEc2ModelPrefixList) WithPrefixListName(a string) *ServicesEc2ModelPrefixList {
	conv_a := javabind.NewGoToJavaString()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "withPrefixListName", "com/amazonaws/services/ec2/model/PrefixList", conv_a.Value().Cast("java/lang/String"))
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
	unique_x := &ServicesEc2ModelPrefixList{}
	unique_x.Callable = dst
	return unique_x
}

// public java.util.List<java.lang.String> getCidrs()
func (jbobject *ServicesEc2ModelPrefixList) GetCidrs() []string {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getCidrs", "java/util/List")
	if err != nil {
		panic(err)
	}
	retconv := javabind.NewJavaToGoList(javabind.NewJavaToGoString())
	dst := new([]string)
	retconv.Dest(dst)
	if err := retconv.Convert(javabind.ObjectRef(jret)); err != nil {
		panic(err)
	}
	retconv.CleanUp()
	return *dst
}

// public void setCidrs(java.util.Collection<java.lang.String>)
func (jbobject *ServicesEc2ModelPrefixList) SetCidrs(a []string)  {
	conv_a := javabind.NewGoToJavaCollection(javabind.NewGoToJavaString())
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "setCidrs", javabind.Void, conv_a.Value().Cast("java/util/Collection"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()

}

// public com.amazonaws.services.ec2.model.PrefixList withCidrs(java.lang.String...)
func (jbobject *ServicesEc2ModelPrefixList) WithCidrs(a ...string) *ServicesEc2ModelPrefixList {
	conv_a := javabind.NewGoToJavaObjectArray(javabind.NewGoToJavaString(), "java/lang/String")
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "withCidrs", "com/amazonaws/services/ec2/model/PrefixList", conv_a.Value().Cast("java/lang/String"))
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
	unique_x := &ServicesEc2ModelPrefixList{}
	unique_x.Callable = dst
	return unique_x
}

// public com.amazonaws.services.ec2.model.PrefixList withCidrs(java.util.Collection<java.lang.String>)
func (jbobject *ServicesEc2ModelPrefixList) WithCidrs2(a []string) *ServicesEc2ModelPrefixList {
	conv_a := javabind.NewGoToJavaCollection(javabind.NewGoToJavaString())
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "withCidrs", "com/amazonaws/services/ec2/model/PrefixList", conv_a.Value().Cast("java/util/Collection"))
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
	unique_x := &ServicesEc2ModelPrefixList{}
	unique_x.Callable = dst
	return unique_x
}

// public java.lang.String toString()
func (jbobject *ServicesEc2ModelPrefixList) ToString() string {
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
func (jbobject *ServicesEc2ModelPrefixList) Equals(a interface{}) bool {
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
func (jbobject *ServicesEc2ModelPrefixList) HashCode() int {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "hashCode", javabind.Int)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

// public com.amazonaws.services.ec2.model.PrefixList clone()
func (jbobject *ServicesEc2ModelPrefixList) Clone() *ServicesEc2ModelPrefixList {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "clone", "com/amazonaws/services/ec2/model/PrefixList")
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
	unique_x := &ServicesEc2ModelPrefixList{}
	unique_x.Callable = dst
	return unique_x
}

// public java.lang.Object clone() throws java.lang.CloneNotSupportedException
func (jbobject *ServicesEc2ModelPrefixList) Clone2() (*JavaLangObject, error) {
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


