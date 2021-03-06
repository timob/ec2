package ec2

import "github.com/timob/javabind"

type ServicesEc2ModelCreateReservedInstancesListingResultInterface interface {
	JavaLangObjectInterface

	// public java.util.List<com.amazonaws.services.ec2.model.ReservedInstancesListing> getReservedInstancesListings()
	GetReservedInstancesListings() []*ServicesEc2ModelReservedInstancesListing

	// public void setReservedInstancesListings(java.util.Collection<com.amazonaws.services.ec2.model.ReservedInstancesListing>)
	SetReservedInstancesListings(a []*ServicesEc2ModelReservedInstancesListing) 

	// public com.amazonaws.services.ec2.model.CreateReservedInstancesListingResult withReservedInstancesListings(com.amazonaws.services.ec2.model.ReservedInstancesListing...)
	WithReservedInstancesListings(a ...*ServicesEc2ModelReservedInstancesListing) *ServicesEc2ModelCreateReservedInstancesListingResult

	// public com.amazonaws.services.ec2.model.CreateReservedInstancesListingResult withReservedInstancesListings(java.util.Collection<com.amazonaws.services.ec2.model.ReservedInstancesListing>)
	WithReservedInstancesListings2(a []*ServicesEc2ModelReservedInstancesListing) *ServicesEc2ModelCreateReservedInstancesListingResult

	// public com.amazonaws.services.ec2.model.CreateReservedInstancesListingResult clone()
	Clone() *ServicesEc2ModelCreateReservedInstancesListingResult
}

type ServicesEc2ModelCreateReservedInstancesListingResult struct {
	JavaLangObject
}

// public com.amazonaws.services.ec2.model.CreateReservedInstancesListingResult()
func NewServicesEc2ModelCreateReservedInstancesListingResult() (*ServicesEc2ModelCreateReservedInstancesListingResult) {

	obj, err := javabind.GetEnv().NewObject("com/amazonaws/services/ec2/model/CreateReservedInstancesListingResult")
	if err != nil {
		panic(err)
	}
	x := &ServicesEc2ModelCreateReservedInstancesListingResult{}
	x.Callable = &javabind.Callable{obj}
	return x
}

// public java.util.List<com.amazonaws.services.ec2.model.ReservedInstancesListing> getReservedInstancesListings()
func (jbobject *ServicesEc2ModelCreateReservedInstancesListingResult) GetReservedInstancesListings() []*ServicesEc2ModelReservedInstancesListing {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getReservedInstancesListings", "java/util/List")
	if err != nil {
		panic(err)
	}
	retconv := javabind.NewJavaToGoList(javabind.NewJavaToGoCallable())
	dst := new([]*ServicesEc2ModelReservedInstancesListing)
	retconv.Dest(dst)
	if err := retconv.Convert(javabind.ObjectRef(jret)); err != nil {
		panic(err)
	}
	retconv.CleanUp()
	return *dst
}

// public void setReservedInstancesListings(java.util.Collection<com.amazonaws.services.ec2.model.ReservedInstancesListing>)
func (jbobject *ServicesEc2ModelCreateReservedInstancesListingResult) SetReservedInstancesListings(a []*ServicesEc2ModelReservedInstancesListing)  {
	conv_a := javabind.NewGoToJavaCollection(javabind.NewGoToJavaCallable())
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "setReservedInstancesListings", javabind.Void, conv_a.Value().Cast("java/util/Collection"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()

}

// public com.amazonaws.services.ec2.model.CreateReservedInstancesListingResult withReservedInstancesListings(com.amazonaws.services.ec2.model.ReservedInstancesListing...)
func (jbobject *ServicesEc2ModelCreateReservedInstancesListingResult) WithReservedInstancesListings(a ...*ServicesEc2ModelReservedInstancesListing) *ServicesEc2ModelCreateReservedInstancesListingResult {
	conv_a := javabind.NewGoToJavaObjectArray(javabind.NewGoToJavaCallable(), "com/amazonaws/services/ec2/model/ReservedInstancesListing")
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "withReservedInstancesListings", "com/amazonaws/services/ec2/model/CreateReservedInstancesListingResult", conv_a.Value().Cast("com/amazonaws/services/ec2/model/ReservedInstancesListing"))
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
	unique_x := &ServicesEc2ModelCreateReservedInstancesListingResult{}
	unique_x.Callable = dst
	return unique_x
}

// public com.amazonaws.services.ec2.model.CreateReservedInstancesListingResult withReservedInstancesListings(java.util.Collection<com.amazonaws.services.ec2.model.ReservedInstancesListing>)
func (jbobject *ServicesEc2ModelCreateReservedInstancesListingResult) WithReservedInstancesListings2(a []*ServicesEc2ModelReservedInstancesListing) *ServicesEc2ModelCreateReservedInstancesListingResult {
	conv_a := javabind.NewGoToJavaCollection(javabind.NewGoToJavaCallable())
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "withReservedInstancesListings", "com/amazonaws/services/ec2/model/CreateReservedInstancesListingResult", conv_a.Value().Cast("java/util/Collection"))
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
	unique_x := &ServicesEc2ModelCreateReservedInstancesListingResult{}
	unique_x.Callable = dst
	return unique_x
}

// public java.lang.String toString()
func (jbobject *ServicesEc2ModelCreateReservedInstancesListingResult) ToString() string {
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
func (jbobject *ServicesEc2ModelCreateReservedInstancesListingResult) Equals(a interface{}) bool {
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
func (jbobject *ServicesEc2ModelCreateReservedInstancesListingResult) HashCode() int {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "hashCode", javabind.Int)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

// public com.amazonaws.services.ec2.model.CreateReservedInstancesListingResult clone()
func (jbobject *ServicesEc2ModelCreateReservedInstancesListingResult) Clone() *ServicesEc2ModelCreateReservedInstancesListingResult {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "clone", "com/amazonaws/services/ec2/model/CreateReservedInstancesListingResult")
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
	unique_x := &ServicesEc2ModelCreateReservedInstancesListingResult{}
	unique_x.Callable = dst
	return unique_x
}

// public java.lang.Object clone() throws java.lang.CloneNotSupportedException
func (jbobject *ServicesEc2ModelCreateReservedInstancesListingResult) Clone2() (*JavaLangObject, error) {
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


