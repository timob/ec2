package ec2

import "github.com/timob/javabind"

type ServicesEc2ModelCancelReservedInstancesListingResultInterface interface {
	JavaLangObjectInterface

	// public java.util.List<com.amazonaws.services.ec2.model.ReservedInstancesListing> getReservedInstancesListings()
	GetReservedInstancesListings() []*ServicesEc2ModelReservedInstancesListing

	// public void setReservedInstancesListings(java.util.Collection<com.amazonaws.services.ec2.model.ReservedInstancesListing>)
	SetReservedInstancesListings(a []*ServicesEc2ModelReservedInstancesListing) 

	// public com.amazonaws.services.ec2.model.CancelReservedInstancesListingResult withReservedInstancesListings(com.amazonaws.services.ec2.model.ReservedInstancesListing...)
	WithReservedInstancesListings(a ...*ServicesEc2ModelReservedInstancesListing) *ServicesEc2ModelCancelReservedInstancesListingResult

	// public com.amazonaws.services.ec2.model.CancelReservedInstancesListingResult withReservedInstancesListings(java.util.Collection<com.amazonaws.services.ec2.model.ReservedInstancesListing>)
	WithReservedInstancesListings2(a []*ServicesEc2ModelReservedInstancesListing) *ServicesEc2ModelCancelReservedInstancesListingResult

	// public com.amazonaws.services.ec2.model.CancelReservedInstancesListingResult clone()
	Clone() *ServicesEc2ModelCancelReservedInstancesListingResult
}

type ServicesEc2ModelCancelReservedInstancesListingResult struct {
	JavaLangObject
}

// public com.amazonaws.services.ec2.model.CancelReservedInstancesListingResult()
func NewServicesEc2ModelCancelReservedInstancesListingResult() (*ServicesEc2ModelCancelReservedInstancesListingResult) {

	obj, err := javabind.GetEnv().NewObject("com/amazonaws/services/ec2/model/CancelReservedInstancesListingResult")
	if err != nil {
		panic(err)
	}
	x := &ServicesEc2ModelCancelReservedInstancesListingResult{}
	x.Callable = &javabind.Callable{obj}
	return x
}

// public java.util.List<com.amazonaws.services.ec2.model.ReservedInstancesListing> getReservedInstancesListings()
func (jbobject *ServicesEc2ModelCancelReservedInstancesListingResult) GetReservedInstancesListings() []*ServicesEc2ModelReservedInstancesListing {
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
func (jbobject *ServicesEc2ModelCancelReservedInstancesListingResult) SetReservedInstancesListings(a []*ServicesEc2ModelReservedInstancesListing)  {
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

// public com.amazonaws.services.ec2.model.CancelReservedInstancesListingResult withReservedInstancesListings(com.amazonaws.services.ec2.model.ReservedInstancesListing...)
func (jbobject *ServicesEc2ModelCancelReservedInstancesListingResult) WithReservedInstancesListings(a ...*ServicesEc2ModelReservedInstancesListing) *ServicesEc2ModelCancelReservedInstancesListingResult {
	conv_a := javabind.NewGoToJavaObjectArray(javabind.NewGoToJavaCallable(), "com/amazonaws/services/ec2/model/ReservedInstancesListing")
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "withReservedInstancesListings", "com/amazonaws/services/ec2/model/CancelReservedInstancesListingResult", conv_a.Value().Cast("com/amazonaws/services/ec2/model/ReservedInstancesListing"))
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
	unique_x := &ServicesEc2ModelCancelReservedInstancesListingResult{}
	unique_x.Callable = dst
	return unique_x
}

// public com.amazonaws.services.ec2.model.CancelReservedInstancesListingResult withReservedInstancesListings(java.util.Collection<com.amazonaws.services.ec2.model.ReservedInstancesListing>)
func (jbobject *ServicesEc2ModelCancelReservedInstancesListingResult) WithReservedInstancesListings2(a []*ServicesEc2ModelReservedInstancesListing) *ServicesEc2ModelCancelReservedInstancesListingResult {
	conv_a := javabind.NewGoToJavaCollection(javabind.NewGoToJavaCallable())
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "withReservedInstancesListings", "com/amazonaws/services/ec2/model/CancelReservedInstancesListingResult", conv_a.Value().Cast("java/util/Collection"))
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
	unique_x := &ServicesEc2ModelCancelReservedInstancesListingResult{}
	unique_x.Callable = dst
	return unique_x
}

// public java.lang.String toString()
func (jbobject *ServicesEc2ModelCancelReservedInstancesListingResult) ToString() string {
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
func (jbobject *ServicesEc2ModelCancelReservedInstancesListingResult) Equals(a interface{}) bool {
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
func (jbobject *ServicesEc2ModelCancelReservedInstancesListingResult) HashCode() int {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "hashCode", javabind.Int)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

// public com.amazonaws.services.ec2.model.CancelReservedInstancesListingResult clone()
func (jbobject *ServicesEc2ModelCancelReservedInstancesListingResult) Clone() *ServicesEc2ModelCancelReservedInstancesListingResult {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "clone", "com/amazonaws/services/ec2/model/CancelReservedInstancesListingResult")
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
	unique_x := &ServicesEc2ModelCancelReservedInstancesListingResult{}
	unique_x.Callable = dst
	return unique_x
}

// public java.lang.Object clone() throws java.lang.CloneNotSupportedException
func (jbobject *ServicesEc2ModelCancelReservedInstancesListingResult) Clone2() (*JavaLangObject, error) {
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


