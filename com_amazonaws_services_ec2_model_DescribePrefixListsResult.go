package ec2

import "github.com/timob/javabind"

type ServicesEc2ModelDescribePrefixListsResultInterface interface {
	JavaLangObjectInterface

	// public java.util.List<com.amazonaws.services.ec2.model.PrefixList> getPrefixLists()
	GetPrefixLists() []*ServicesEc2ModelPrefixList

	// public void setPrefixLists(java.util.Collection<com.amazonaws.services.ec2.model.PrefixList>)
	SetPrefixLists(a []*ServicesEc2ModelPrefixList) 

	// public com.amazonaws.services.ec2.model.DescribePrefixListsResult withPrefixLists(com.amazonaws.services.ec2.model.PrefixList...)
	WithPrefixLists(a ...*ServicesEc2ModelPrefixList) *ServicesEc2ModelDescribePrefixListsResult

	// public com.amazonaws.services.ec2.model.DescribePrefixListsResult withPrefixLists(java.util.Collection<com.amazonaws.services.ec2.model.PrefixList>)
	WithPrefixLists2(a []*ServicesEc2ModelPrefixList) *ServicesEc2ModelDescribePrefixListsResult

	// public void setNextToken(java.lang.String)
	SetNextToken(a string) 

	// public java.lang.String getNextToken()
	GetNextToken() string

	// public com.amazonaws.services.ec2.model.DescribePrefixListsResult withNextToken(java.lang.String)
	WithNextToken(a string) *ServicesEc2ModelDescribePrefixListsResult

	// public com.amazonaws.services.ec2.model.DescribePrefixListsResult clone()
	Clone() *ServicesEc2ModelDescribePrefixListsResult
}

type ServicesEc2ModelDescribePrefixListsResult struct {
	JavaLangObject
}

// public com.amazonaws.services.ec2.model.DescribePrefixListsResult()
func NewServicesEc2ModelDescribePrefixListsResult() (*ServicesEc2ModelDescribePrefixListsResult) {

	obj, err := javabind.GetEnv().NewObject("com/amazonaws/services/ec2/model/DescribePrefixListsResult")
	if err != nil {
		panic(err)
	}
	x := &ServicesEc2ModelDescribePrefixListsResult{}
	x.Callable = &javabind.Callable{obj}
	return x
}

// public java.util.List<com.amazonaws.services.ec2.model.PrefixList> getPrefixLists()
func (jbobject *ServicesEc2ModelDescribePrefixListsResult) GetPrefixLists() []*ServicesEc2ModelPrefixList {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getPrefixLists", "java/util/List")
	if err != nil {
		panic(err)
	}
	retconv := javabind.NewJavaToGoList(javabind.NewJavaToGoCallable())
	dst := new([]*ServicesEc2ModelPrefixList)
	retconv.Dest(dst)
	if err := retconv.Convert(javabind.ObjectRef(jret)); err != nil {
		panic(err)
	}
	retconv.CleanUp()
	return *dst
}

// public void setPrefixLists(java.util.Collection<com.amazonaws.services.ec2.model.PrefixList>)
func (jbobject *ServicesEc2ModelDescribePrefixListsResult) SetPrefixLists(a []*ServicesEc2ModelPrefixList)  {
	conv_a := javabind.NewGoToJavaCollection(javabind.NewGoToJavaCallable())
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "setPrefixLists", javabind.Void, conv_a.Value().Cast("java/util/Collection"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()

}

// public com.amazonaws.services.ec2.model.DescribePrefixListsResult withPrefixLists(com.amazonaws.services.ec2.model.PrefixList...)
func (jbobject *ServicesEc2ModelDescribePrefixListsResult) WithPrefixLists(a ...*ServicesEc2ModelPrefixList) *ServicesEc2ModelDescribePrefixListsResult {
	conv_a := javabind.NewGoToJavaObjectArray(javabind.NewGoToJavaCallable(), "com/amazonaws/services/ec2/model/PrefixList")
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "withPrefixLists", "com/amazonaws/services/ec2/model/DescribePrefixListsResult", conv_a.Value().Cast("com/amazonaws/services/ec2/model/PrefixList"))
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
	unique_x := &ServicesEc2ModelDescribePrefixListsResult{}
	unique_x.Callable = dst
	return unique_x
}

// public com.amazonaws.services.ec2.model.DescribePrefixListsResult withPrefixLists(java.util.Collection<com.amazonaws.services.ec2.model.PrefixList>)
func (jbobject *ServicesEc2ModelDescribePrefixListsResult) WithPrefixLists2(a []*ServicesEc2ModelPrefixList) *ServicesEc2ModelDescribePrefixListsResult {
	conv_a := javabind.NewGoToJavaCollection(javabind.NewGoToJavaCallable())
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "withPrefixLists", "com/amazonaws/services/ec2/model/DescribePrefixListsResult", conv_a.Value().Cast("java/util/Collection"))
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
	unique_x := &ServicesEc2ModelDescribePrefixListsResult{}
	unique_x.Callable = dst
	return unique_x
}

// public void setNextToken(java.lang.String)
func (jbobject *ServicesEc2ModelDescribePrefixListsResult) SetNextToken(a string)  {
	conv_a := javabind.NewGoToJavaString()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "setNextToken", javabind.Void, conv_a.Value().Cast("java/lang/String"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()

}

// public java.lang.String getNextToken()
func (jbobject *ServicesEc2ModelDescribePrefixListsResult) GetNextToken() string {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getNextToken", "java/lang/String")
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

// public com.amazonaws.services.ec2.model.DescribePrefixListsResult withNextToken(java.lang.String)
func (jbobject *ServicesEc2ModelDescribePrefixListsResult) WithNextToken(a string) *ServicesEc2ModelDescribePrefixListsResult {
	conv_a := javabind.NewGoToJavaString()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "withNextToken", "com/amazonaws/services/ec2/model/DescribePrefixListsResult", conv_a.Value().Cast("java/lang/String"))
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
	unique_x := &ServicesEc2ModelDescribePrefixListsResult{}
	unique_x.Callable = dst
	return unique_x
}

// public java.lang.String toString()
func (jbobject *ServicesEc2ModelDescribePrefixListsResult) ToString() string {
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
func (jbobject *ServicesEc2ModelDescribePrefixListsResult) Equals(a interface{}) bool {
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
func (jbobject *ServicesEc2ModelDescribePrefixListsResult) HashCode() int {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "hashCode", javabind.Int)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

// public com.amazonaws.services.ec2.model.DescribePrefixListsResult clone()
func (jbobject *ServicesEc2ModelDescribePrefixListsResult) Clone() *ServicesEc2ModelDescribePrefixListsResult {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "clone", "com/amazonaws/services/ec2/model/DescribePrefixListsResult")
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
	unique_x := &ServicesEc2ModelDescribePrefixListsResult{}
	unique_x.Callable = dst
	return unique_x
}

// public java.lang.Object clone() throws java.lang.CloneNotSupportedException
func (jbobject *ServicesEc2ModelDescribePrefixListsResult) Clone2() (*JavaLangObject, error) {
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


