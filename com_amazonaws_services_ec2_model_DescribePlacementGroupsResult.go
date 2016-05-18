package ec2

import "github.com/timob/javabind"

type ServicesEc2ModelDescribePlacementGroupsResultInterface interface {
	JavaLangObjectInterface

	// public java.util.List<com.amazonaws.services.ec2.model.PlacementGroup> getPlacementGroups()
	GetPlacementGroups() []*ServicesEc2ModelPlacementGroup

	// public void setPlacementGroups(java.util.Collection<com.amazonaws.services.ec2.model.PlacementGroup>)
	SetPlacementGroups(a []*ServicesEc2ModelPlacementGroup) 

	// public com.amazonaws.services.ec2.model.DescribePlacementGroupsResult withPlacementGroups(com.amazonaws.services.ec2.model.PlacementGroup...)
	WithPlacementGroups(a ...*ServicesEc2ModelPlacementGroup) *ServicesEc2ModelDescribePlacementGroupsResult

	// public com.amazonaws.services.ec2.model.DescribePlacementGroupsResult withPlacementGroups(java.util.Collection<com.amazonaws.services.ec2.model.PlacementGroup>)
	WithPlacementGroups2(a []*ServicesEc2ModelPlacementGroup) *ServicesEc2ModelDescribePlacementGroupsResult

	// public com.amazonaws.services.ec2.model.DescribePlacementGroupsResult clone()
	Clone() *ServicesEc2ModelDescribePlacementGroupsResult
}

type ServicesEc2ModelDescribePlacementGroupsResult struct {
	JavaLangObject
}

// public com.amazonaws.services.ec2.model.DescribePlacementGroupsResult()
func NewServicesEc2ModelDescribePlacementGroupsResult() (*ServicesEc2ModelDescribePlacementGroupsResult) {

	obj, err := javabind.GetEnv().NewObject("com/amazonaws/services/ec2/model/DescribePlacementGroupsResult")
	if err != nil {
		panic(err)
	}
	x := &ServicesEc2ModelDescribePlacementGroupsResult{}
	x.Callable = &javabind.Callable{obj}
	return x
}

// public java.util.List<com.amazonaws.services.ec2.model.PlacementGroup> getPlacementGroups()
func (jbobject *ServicesEc2ModelDescribePlacementGroupsResult) GetPlacementGroups() []*ServicesEc2ModelPlacementGroup {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getPlacementGroups", "java/util/List")
	if err != nil {
		panic(err)
	}
	retconv := javabind.NewJavaToGoList(javabind.NewJavaToGoCallable())
	dst := new([]*ServicesEc2ModelPlacementGroup)
	retconv.Dest(dst)
	if err := retconv.Convert(javabind.ObjectRef(jret)); err != nil {
		panic(err)
	}
	retconv.CleanUp()
	return *dst
}

// public void setPlacementGroups(java.util.Collection<com.amazonaws.services.ec2.model.PlacementGroup>)
func (jbobject *ServicesEc2ModelDescribePlacementGroupsResult) SetPlacementGroups(a []*ServicesEc2ModelPlacementGroup)  {
	conv_a := javabind.NewGoToJavaCollection(javabind.NewGoToJavaCallable())
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "setPlacementGroups", javabind.Void, conv_a.Value().Cast("java/util/Collection"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()

}

// public com.amazonaws.services.ec2.model.DescribePlacementGroupsResult withPlacementGroups(com.amazonaws.services.ec2.model.PlacementGroup...)
func (jbobject *ServicesEc2ModelDescribePlacementGroupsResult) WithPlacementGroups(a ...*ServicesEc2ModelPlacementGroup) *ServicesEc2ModelDescribePlacementGroupsResult {
	conv_a := javabind.NewGoToJavaObjectArray(javabind.NewGoToJavaCallable(), "com/amazonaws/services/ec2/model/PlacementGroup")
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "withPlacementGroups", "com/amazonaws/services/ec2/model/DescribePlacementGroupsResult", conv_a.Value().Cast("com/amazonaws/services/ec2/model/PlacementGroup"))
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
	unique_x := &ServicesEc2ModelDescribePlacementGroupsResult{}
	unique_x.Callable = dst
	return unique_x
}

// public com.amazonaws.services.ec2.model.DescribePlacementGroupsResult withPlacementGroups(java.util.Collection<com.amazonaws.services.ec2.model.PlacementGroup>)
func (jbobject *ServicesEc2ModelDescribePlacementGroupsResult) WithPlacementGroups2(a []*ServicesEc2ModelPlacementGroup) *ServicesEc2ModelDescribePlacementGroupsResult {
	conv_a := javabind.NewGoToJavaCollection(javabind.NewGoToJavaCallable())
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "withPlacementGroups", "com/amazonaws/services/ec2/model/DescribePlacementGroupsResult", conv_a.Value().Cast("java/util/Collection"))
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
	unique_x := &ServicesEc2ModelDescribePlacementGroupsResult{}
	unique_x.Callable = dst
	return unique_x
}

// public java.lang.String toString()
func (jbobject *ServicesEc2ModelDescribePlacementGroupsResult) ToString() string {
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
func (jbobject *ServicesEc2ModelDescribePlacementGroupsResult) Equals(a interface{}) bool {
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
func (jbobject *ServicesEc2ModelDescribePlacementGroupsResult) HashCode() int {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "hashCode", javabind.Int)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

// public com.amazonaws.services.ec2.model.DescribePlacementGroupsResult clone()
func (jbobject *ServicesEc2ModelDescribePlacementGroupsResult) Clone() *ServicesEc2ModelDescribePlacementGroupsResult {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "clone", "com/amazonaws/services/ec2/model/DescribePlacementGroupsResult")
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
	unique_x := &ServicesEc2ModelDescribePlacementGroupsResult{}
	unique_x.Callable = dst
	return unique_x
}

// public java.lang.Object clone() throws java.lang.CloneNotSupportedException
func (jbobject *ServicesEc2ModelDescribePlacementGroupsResult) Clone2() (*JavaLangObject, error) {
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


