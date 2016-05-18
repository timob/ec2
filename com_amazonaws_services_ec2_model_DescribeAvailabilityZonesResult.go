package ec2

import "github.com/timob/javabind"

type ServicesEc2ModelDescribeAvailabilityZonesResultInterface interface {
	JavaLangObjectInterface

	// public java.util.List<com.amazonaws.services.ec2.model.AvailabilityZone> getAvailabilityZones()
	GetAvailabilityZones() []*ServicesEc2ModelAvailabilityZone

	// public void setAvailabilityZones(java.util.Collection<com.amazonaws.services.ec2.model.AvailabilityZone>)
	SetAvailabilityZones(a []*ServicesEc2ModelAvailabilityZone) 

	// public com.amazonaws.services.ec2.model.DescribeAvailabilityZonesResult withAvailabilityZones(com.amazonaws.services.ec2.model.AvailabilityZone...)
	WithAvailabilityZones(a ...*ServicesEc2ModelAvailabilityZone) *ServicesEc2ModelDescribeAvailabilityZonesResult

	// public com.amazonaws.services.ec2.model.DescribeAvailabilityZonesResult withAvailabilityZones(java.util.Collection<com.amazonaws.services.ec2.model.AvailabilityZone>)
	WithAvailabilityZones2(a []*ServicesEc2ModelAvailabilityZone) *ServicesEc2ModelDescribeAvailabilityZonesResult

	// public com.amazonaws.services.ec2.model.DescribeAvailabilityZonesResult clone()
	Clone() *ServicesEc2ModelDescribeAvailabilityZonesResult
}

type ServicesEc2ModelDescribeAvailabilityZonesResult struct {
	JavaLangObject
}

// public com.amazonaws.services.ec2.model.DescribeAvailabilityZonesResult()
func NewServicesEc2ModelDescribeAvailabilityZonesResult() (*ServicesEc2ModelDescribeAvailabilityZonesResult) {

	obj, err := javabind.GetEnv().NewObject("com/amazonaws/services/ec2/model/DescribeAvailabilityZonesResult")
	if err != nil {
		panic(err)
	}
	x := &ServicesEc2ModelDescribeAvailabilityZonesResult{}
	x.Callable = &javabind.Callable{obj}
	return x
}

// public java.util.List<com.amazonaws.services.ec2.model.AvailabilityZone> getAvailabilityZones()
func (jbobject *ServicesEc2ModelDescribeAvailabilityZonesResult) GetAvailabilityZones() []*ServicesEc2ModelAvailabilityZone {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getAvailabilityZones", "java/util/List")
	if err != nil {
		panic(err)
	}
	retconv := javabind.NewJavaToGoList(javabind.NewJavaToGoCallable())
	dst := new([]*ServicesEc2ModelAvailabilityZone)
	retconv.Dest(dst)
	if err := retconv.Convert(javabind.ObjectRef(jret)); err != nil {
		panic(err)
	}
	retconv.CleanUp()
	return *dst
}

// public void setAvailabilityZones(java.util.Collection<com.amazonaws.services.ec2.model.AvailabilityZone>)
func (jbobject *ServicesEc2ModelDescribeAvailabilityZonesResult) SetAvailabilityZones(a []*ServicesEc2ModelAvailabilityZone)  {
	conv_a := javabind.NewGoToJavaCollection(javabind.NewGoToJavaCallable())
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "setAvailabilityZones", javabind.Void, conv_a.Value().Cast("java/util/Collection"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()

}

// public com.amazonaws.services.ec2.model.DescribeAvailabilityZonesResult withAvailabilityZones(com.amazonaws.services.ec2.model.AvailabilityZone...)
func (jbobject *ServicesEc2ModelDescribeAvailabilityZonesResult) WithAvailabilityZones(a ...*ServicesEc2ModelAvailabilityZone) *ServicesEc2ModelDescribeAvailabilityZonesResult {
	conv_a := javabind.NewGoToJavaObjectArray(javabind.NewGoToJavaCallable(), "com/amazonaws/services/ec2/model/AvailabilityZone")
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "withAvailabilityZones", "com/amazonaws/services/ec2/model/DescribeAvailabilityZonesResult", conv_a.Value().Cast("com/amazonaws/services/ec2/model/AvailabilityZone"))
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
	unique_x := &ServicesEc2ModelDescribeAvailabilityZonesResult{}
	unique_x.Callable = dst
	return unique_x
}

// public com.amazonaws.services.ec2.model.DescribeAvailabilityZonesResult withAvailabilityZones(java.util.Collection<com.amazonaws.services.ec2.model.AvailabilityZone>)
func (jbobject *ServicesEc2ModelDescribeAvailabilityZonesResult) WithAvailabilityZones2(a []*ServicesEc2ModelAvailabilityZone) *ServicesEc2ModelDescribeAvailabilityZonesResult {
	conv_a := javabind.NewGoToJavaCollection(javabind.NewGoToJavaCallable())
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "withAvailabilityZones", "com/amazonaws/services/ec2/model/DescribeAvailabilityZonesResult", conv_a.Value().Cast("java/util/Collection"))
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
	unique_x := &ServicesEc2ModelDescribeAvailabilityZonesResult{}
	unique_x.Callable = dst
	return unique_x
}

// public java.lang.String toString()
func (jbobject *ServicesEc2ModelDescribeAvailabilityZonesResult) ToString() string {
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
func (jbobject *ServicesEc2ModelDescribeAvailabilityZonesResult) Equals(a interface{}) bool {
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
func (jbobject *ServicesEc2ModelDescribeAvailabilityZonesResult) HashCode() int {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "hashCode", javabind.Int)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

// public com.amazonaws.services.ec2.model.DescribeAvailabilityZonesResult clone()
func (jbobject *ServicesEc2ModelDescribeAvailabilityZonesResult) Clone() *ServicesEc2ModelDescribeAvailabilityZonesResult {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "clone", "com/amazonaws/services/ec2/model/DescribeAvailabilityZonesResult")
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
	unique_x := &ServicesEc2ModelDescribeAvailabilityZonesResult{}
	unique_x.Callable = dst
	return unique_x
}

// public java.lang.Object clone() throws java.lang.CloneNotSupportedException
func (jbobject *ServicesEc2ModelDescribeAvailabilityZonesResult) Clone2() (*JavaLangObject, error) {
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


