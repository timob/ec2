package ec2

import "github.com/timob/javabind"

type ServicesEc2ModelDescribeReservedInstancesOfferingsRequestInterface interface {
	AmazonWebServiceRequestInterface

	// public java.util.List<java.lang.String> getReservedInstancesOfferingIds()
	GetReservedInstancesOfferingIds() []string

	// public void setReservedInstancesOfferingIds(java.util.Collection<java.lang.String>)
	SetReservedInstancesOfferingIds(a []string) 

	// public com.amazonaws.services.ec2.model.DescribeReservedInstancesOfferingsRequest withReservedInstancesOfferingIds(java.lang.String...)
	WithReservedInstancesOfferingIds(a ...string) *ServicesEc2ModelDescribeReservedInstancesOfferingsRequest

	// public com.amazonaws.services.ec2.model.DescribeReservedInstancesOfferingsRequest withReservedInstancesOfferingIds(java.util.Collection<java.lang.String>)
	WithReservedInstancesOfferingIds2(a []string) *ServicesEc2ModelDescribeReservedInstancesOfferingsRequest

	// public void setInstanceType(java.lang.String)
	SetInstanceType2(a string) 

	// public java.lang.String getInstanceType()
	GetInstanceType() string

	// public com.amazonaws.services.ec2.model.DescribeReservedInstancesOfferingsRequest withInstanceType(java.lang.String)
	WithInstanceType2(a string) *ServicesEc2ModelDescribeReservedInstancesOfferingsRequest

	// public void setInstanceType(com.amazonaws.services.ec2.model.InstanceType)
	SetInstanceType(a ServicesEc2ModelInstanceTypeInterface) 

	// public com.amazonaws.services.ec2.model.DescribeReservedInstancesOfferingsRequest withInstanceType(com.amazonaws.services.ec2.model.InstanceType)
	WithInstanceType(a ServicesEc2ModelInstanceTypeInterface) *ServicesEc2ModelDescribeReservedInstancesOfferingsRequest

	// public void setAvailabilityZone(java.lang.String)
	SetAvailabilityZone(a string) 

	// public java.lang.String getAvailabilityZone()
	GetAvailabilityZone() string

	// public com.amazonaws.services.ec2.model.DescribeReservedInstancesOfferingsRequest withAvailabilityZone(java.lang.String)
	WithAvailabilityZone(a string) *ServicesEc2ModelDescribeReservedInstancesOfferingsRequest

	// public void setProductDescription(java.lang.String)
	SetProductDescription2(a string) 

	// public java.lang.String getProductDescription()
	GetProductDescription() string

	// public com.amazonaws.services.ec2.model.DescribeReservedInstancesOfferingsRequest withProductDescription(java.lang.String)
	WithProductDescription2(a string) *ServicesEc2ModelDescribeReservedInstancesOfferingsRequest

	// public void setProductDescription(com.amazonaws.services.ec2.model.RIProductDescription)
	SetProductDescription(a ServicesEc2ModelRIProductDescriptionInterface) 

	// public com.amazonaws.services.ec2.model.DescribeReservedInstancesOfferingsRequest withProductDescription(com.amazonaws.services.ec2.model.RIProductDescription)
	WithProductDescription(a ServicesEc2ModelRIProductDescriptionInterface) *ServicesEc2ModelDescribeReservedInstancesOfferingsRequest

	// public java.util.List<com.amazonaws.services.ec2.model.Filter> getFilters()
	GetFilters() []*ServicesEc2ModelFilter

	// public void setFilters(java.util.Collection<com.amazonaws.services.ec2.model.Filter>)
	SetFilters(a []*ServicesEc2ModelFilter) 

	// public com.amazonaws.services.ec2.model.DescribeReservedInstancesOfferingsRequest withFilters(com.amazonaws.services.ec2.model.Filter...)
	WithFilters(a ...*ServicesEc2ModelFilter) *ServicesEc2ModelDescribeReservedInstancesOfferingsRequest

	// public com.amazonaws.services.ec2.model.DescribeReservedInstancesOfferingsRequest withFilters(java.util.Collection<com.amazonaws.services.ec2.model.Filter>)
	WithFilters2(a []*ServicesEc2ModelFilter) *ServicesEc2ModelDescribeReservedInstancesOfferingsRequest

	// public void setInstanceTenancy(java.lang.String)
	SetInstanceTenancy2(a string) 

	// public java.lang.String getInstanceTenancy()
	GetInstanceTenancy() string

	// public com.amazonaws.services.ec2.model.DescribeReservedInstancesOfferingsRequest withInstanceTenancy(java.lang.String)
	WithInstanceTenancy2(a string) *ServicesEc2ModelDescribeReservedInstancesOfferingsRequest

	// public void setInstanceTenancy(com.amazonaws.services.ec2.model.Tenancy)
	SetInstanceTenancy(a ServicesEc2ModelTenancyInterface) 

	// public com.amazonaws.services.ec2.model.DescribeReservedInstancesOfferingsRequest withInstanceTenancy(com.amazonaws.services.ec2.model.Tenancy)
	WithInstanceTenancy(a ServicesEc2ModelTenancyInterface) *ServicesEc2ModelDescribeReservedInstancesOfferingsRequest

	// public void setOfferingType(java.lang.String)
	SetOfferingType2(a string) 

	// public java.lang.String getOfferingType()
	GetOfferingType() string

	// public com.amazonaws.services.ec2.model.DescribeReservedInstancesOfferingsRequest withOfferingType(java.lang.String)
	WithOfferingType2(a string) *ServicesEc2ModelDescribeReservedInstancesOfferingsRequest

	// public void setOfferingType(com.amazonaws.services.ec2.model.OfferingTypeValues)
	SetOfferingType(a ServicesEc2ModelOfferingTypeValuesInterface) 

	// public com.amazonaws.services.ec2.model.DescribeReservedInstancesOfferingsRequest withOfferingType(com.amazonaws.services.ec2.model.OfferingTypeValues)
	WithOfferingType(a ServicesEc2ModelOfferingTypeValuesInterface) *ServicesEc2ModelDescribeReservedInstancesOfferingsRequest

	// public void setNextToken(java.lang.String)
	SetNextToken(a string) 

	// public java.lang.String getNextToken()
	GetNextToken() string

	// public com.amazonaws.services.ec2.model.DescribeReservedInstancesOfferingsRequest withNextToken(java.lang.String)
	WithNextToken(a string) *ServicesEc2ModelDescribeReservedInstancesOfferingsRequest

	// public void setMaxResults(java.lang.Integer)
	SetMaxResults(a int) 

	// public java.lang.Integer getMaxResults()
	GetMaxResults() int

	// public com.amazonaws.services.ec2.model.DescribeReservedInstancesOfferingsRequest withMaxResults(java.lang.Integer)
	WithMaxResults(a int) *ServicesEc2ModelDescribeReservedInstancesOfferingsRequest

	// public void setIncludeMarketplace(java.lang.Boolean)
	SetIncludeMarketplace(a bool) 

	// public java.lang.Boolean getIncludeMarketplace()
	GetIncludeMarketplace() bool

	// public com.amazonaws.services.ec2.model.DescribeReservedInstancesOfferingsRequest withIncludeMarketplace(java.lang.Boolean)
	WithIncludeMarketplace(a bool) *ServicesEc2ModelDescribeReservedInstancesOfferingsRequest

	// public java.lang.Boolean isIncludeMarketplace()
	IsIncludeMarketplace() bool

	// public void setMinDuration(java.lang.Long)
	SetMinDuration(a int64) 

	// public java.lang.Long getMinDuration()
	GetMinDuration() int64

	// public com.amazonaws.services.ec2.model.DescribeReservedInstancesOfferingsRequest withMinDuration(java.lang.Long)
	WithMinDuration(a int64) *ServicesEc2ModelDescribeReservedInstancesOfferingsRequest

	// public void setMaxDuration(java.lang.Long)
	SetMaxDuration(a int64) 

	// public java.lang.Long getMaxDuration()
	GetMaxDuration() int64

	// public com.amazonaws.services.ec2.model.DescribeReservedInstancesOfferingsRequest withMaxDuration(java.lang.Long)
	WithMaxDuration(a int64) *ServicesEc2ModelDescribeReservedInstancesOfferingsRequest

	// public void setMaxInstanceCount(java.lang.Integer)
	SetMaxInstanceCount(a int) 

	// public java.lang.Integer getMaxInstanceCount()
	GetMaxInstanceCount() int

	// public com.amazonaws.services.ec2.model.DescribeReservedInstancesOfferingsRequest withMaxInstanceCount(java.lang.Integer)
	WithMaxInstanceCount(a int) *ServicesEc2ModelDescribeReservedInstancesOfferingsRequest

	// public com.amazonaws.Request<com.amazonaws.services.ec2.model.DescribeReservedInstancesOfferingsRequest> getDryRunRequest()
	GetDryRunRequest() *Request

	// public com.amazonaws.services.ec2.model.DescribeReservedInstancesOfferingsRequest clone()
	Clone3() *ServicesEc2ModelDescribeReservedInstancesOfferingsRequest
}

type ServicesEc2ModelDescribeReservedInstancesOfferingsRequest struct {
	AmazonWebServiceRequest
}

// public com.amazonaws.services.ec2.model.DescribeReservedInstancesOfferingsRequest()
func NewServicesEc2ModelDescribeReservedInstancesOfferingsRequest() (*ServicesEc2ModelDescribeReservedInstancesOfferingsRequest) {

	obj, err := javabind.GetEnv().NewObject("com/amazonaws/services/ec2/model/DescribeReservedInstancesOfferingsRequest")
	if err != nil {
		panic(err)
	}
	x := &ServicesEc2ModelDescribeReservedInstancesOfferingsRequest{}
	x.Callable = &javabind.Callable{obj}
	return x
}

// public java.util.List<java.lang.String> getReservedInstancesOfferingIds()
func (jbobject *ServicesEc2ModelDescribeReservedInstancesOfferingsRequest) GetReservedInstancesOfferingIds() []string {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getReservedInstancesOfferingIds", "java/util/List")
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

// public void setReservedInstancesOfferingIds(java.util.Collection<java.lang.String>)
func (jbobject *ServicesEc2ModelDescribeReservedInstancesOfferingsRequest) SetReservedInstancesOfferingIds(a []string)  {
	conv_a := javabind.NewGoToJavaCollection(javabind.NewGoToJavaString())
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "setReservedInstancesOfferingIds", javabind.Void, conv_a.Value().Cast("java/util/Collection"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()

}

// public com.amazonaws.services.ec2.model.DescribeReservedInstancesOfferingsRequest withReservedInstancesOfferingIds(java.lang.String...)
func (jbobject *ServicesEc2ModelDescribeReservedInstancesOfferingsRequest) WithReservedInstancesOfferingIds(a ...string) *ServicesEc2ModelDescribeReservedInstancesOfferingsRequest {
	conv_a := javabind.NewGoToJavaObjectArray(javabind.NewGoToJavaString(), "java/lang/String")
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "withReservedInstancesOfferingIds", "com/amazonaws/services/ec2/model/DescribeReservedInstancesOfferingsRequest", conv_a.Value().Cast("java/lang/String"))
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
	unique_x := &ServicesEc2ModelDescribeReservedInstancesOfferingsRequest{}
	unique_x.Callable = dst
	return unique_x
}

// public com.amazonaws.services.ec2.model.DescribeReservedInstancesOfferingsRequest withReservedInstancesOfferingIds(java.util.Collection<java.lang.String>)
func (jbobject *ServicesEc2ModelDescribeReservedInstancesOfferingsRequest) WithReservedInstancesOfferingIds2(a []string) *ServicesEc2ModelDescribeReservedInstancesOfferingsRequest {
	conv_a := javabind.NewGoToJavaCollection(javabind.NewGoToJavaString())
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "withReservedInstancesOfferingIds", "com/amazonaws/services/ec2/model/DescribeReservedInstancesOfferingsRequest", conv_a.Value().Cast("java/util/Collection"))
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
	unique_x := &ServicesEc2ModelDescribeReservedInstancesOfferingsRequest{}
	unique_x.Callable = dst
	return unique_x
}

// public void setInstanceType(java.lang.String)
func (jbobject *ServicesEc2ModelDescribeReservedInstancesOfferingsRequest) SetInstanceType2(a string)  {
	conv_a := javabind.NewGoToJavaString()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "setInstanceType", javabind.Void, conv_a.Value().Cast("java/lang/String"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()

}

// public java.lang.String getInstanceType()
func (jbobject *ServicesEc2ModelDescribeReservedInstancesOfferingsRequest) GetInstanceType() string {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getInstanceType", "java/lang/String")
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

// public com.amazonaws.services.ec2.model.DescribeReservedInstancesOfferingsRequest withInstanceType(java.lang.String)
func (jbobject *ServicesEc2ModelDescribeReservedInstancesOfferingsRequest) WithInstanceType2(a string) *ServicesEc2ModelDescribeReservedInstancesOfferingsRequest {
	conv_a := javabind.NewGoToJavaString()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "withInstanceType", "com/amazonaws/services/ec2/model/DescribeReservedInstancesOfferingsRequest", conv_a.Value().Cast("java/lang/String"))
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
	unique_x := &ServicesEc2ModelDescribeReservedInstancesOfferingsRequest{}
	unique_x.Callable = dst
	return unique_x
}

// public void setInstanceType(com.amazonaws.services.ec2.model.InstanceType)
func (jbobject *ServicesEc2ModelDescribeReservedInstancesOfferingsRequest) SetInstanceType(a ServicesEc2ModelInstanceTypeInterface)  {
	conv_a := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "setInstanceType", javabind.Void, conv_a.Value().Cast("com/amazonaws/services/ec2/model/InstanceType"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()

}

// public com.amazonaws.services.ec2.model.DescribeReservedInstancesOfferingsRequest withInstanceType(com.amazonaws.services.ec2.model.InstanceType)
func (jbobject *ServicesEc2ModelDescribeReservedInstancesOfferingsRequest) WithInstanceType(a ServicesEc2ModelInstanceTypeInterface) *ServicesEc2ModelDescribeReservedInstancesOfferingsRequest {
	conv_a := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "withInstanceType", "com/amazonaws/services/ec2/model/DescribeReservedInstancesOfferingsRequest", conv_a.Value().Cast("com/amazonaws/services/ec2/model/InstanceType"))
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
	unique_x := &ServicesEc2ModelDescribeReservedInstancesOfferingsRequest{}
	unique_x.Callable = dst
	return unique_x
}

// public void setAvailabilityZone(java.lang.String)
func (jbobject *ServicesEc2ModelDescribeReservedInstancesOfferingsRequest) SetAvailabilityZone(a string)  {
	conv_a := javabind.NewGoToJavaString()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "setAvailabilityZone", javabind.Void, conv_a.Value().Cast("java/lang/String"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()

}

// public java.lang.String getAvailabilityZone()
func (jbobject *ServicesEc2ModelDescribeReservedInstancesOfferingsRequest) GetAvailabilityZone() string {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getAvailabilityZone", "java/lang/String")
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

// public com.amazonaws.services.ec2.model.DescribeReservedInstancesOfferingsRequest withAvailabilityZone(java.lang.String)
func (jbobject *ServicesEc2ModelDescribeReservedInstancesOfferingsRequest) WithAvailabilityZone(a string) *ServicesEc2ModelDescribeReservedInstancesOfferingsRequest {
	conv_a := javabind.NewGoToJavaString()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "withAvailabilityZone", "com/amazonaws/services/ec2/model/DescribeReservedInstancesOfferingsRequest", conv_a.Value().Cast("java/lang/String"))
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
	unique_x := &ServicesEc2ModelDescribeReservedInstancesOfferingsRequest{}
	unique_x.Callable = dst
	return unique_x
}

// public void setProductDescription(java.lang.String)
func (jbobject *ServicesEc2ModelDescribeReservedInstancesOfferingsRequest) SetProductDescription2(a string)  {
	conv_a := javabind.NewGoToJavaString()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "setProductDescription", javabind.Void, conv_a.Value().Cast("java/lang/String"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()

}

// public java.lang.String getProductDescription()
func (jbobject *ServicesEc2ModelDescribeReservedInstancesOfferingsRequest) GetProductDescription() string {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getProductDescription", "java/lang/String")
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

// public com.amazonaws.services.ec2.model.DescribeReservedInstancesOfferingsRequest withProductDescription(java.lang.String)
func (jbobject *ServicesEc2ModelDescribeReservedInstancesOfferingsRequest) WithProductDescription2(a string) *ServicesEc2ModelDescribeReservedInstancesOfferingsRequest {
	conv_a := javabind.NewGoToJavaString()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "withProductDescription", "com/amazonaws/services/ec2/model/DescribeReservedInstancesOfferingsRequest", conv_a.Value().Cast("java/lang/String"))
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
	unique_x := &ServicesEc2ModelDescribeReservedInstancesOfferingsRequest{}
	unique_x.Callable = dst
	return unique_x
}

// public void setProductDescription(com.amazonaws.services.ec2.model.RIProductDescription)
func (jbobject *ServicesEc2ModelDescribeReservedInstancesOfferingsRequest) SetProductDescription(a ServicesEc2ModelRIProductDescriptionInterface)  {
	conv_a := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "setProductDescription", javabind.Void, conv_a.Value().Cast("com/amazonaws/services/ec2/model/RIProductDescription"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()

}

// public com.amazonaws.services.ec2.model.DescribeReservedInstancesOfferingsRequest withProductDescription(com.amazonaws.services.ec2.model.RIProductDescription)
func (jbobject *ServicesEc2ModelDescribeReservedInstancesOfferingsRequest) WithProductDescription(a ServicesEc2ModelRIProductDescriptionInterface) *ServicesEc2ModelDescribeReservedInstancesOfferingsRequest {
	conv_a := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "withProductDescription", "com/amazonaws/services/ec2/model/DescribeReservedInstancesOfferingsRequest", conv_a.Value().Cast("com/amazonaws/services/ec2/model/RIProductDescription"))
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
	unique_x := &ServicesEc2ModelDescribeReservedInstancesOfferingsRequest{}
	unique_x.Callable = dst
	return unique_x
}

// public java.util.List<com.amazonaws.services.ec2.model.Filter> getFilters()
func (jbobject *ServicesEc2ModelDescribeReservedInstancesOfferingsRequest) GetFilters() []*ServicesEc2ModelFilter {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getFilters", "java/util/List")
	if err != nil {
		panic(err)
	}
	retconv := javabind.NewJavaToGoList(javabind.NewJavaToGoCallable())
	dst := new([]*ServicesEc2ModelFilter)
	retconv.Dest(dst)
	if err := retconv.Convert(javabind.ObjectRef(jret)); err != nil {
		panic(err)
	}
	retconv.CleanUp()
	return *dst
}

// public void setFilters(java.util.Collection<com.amazonaws.services.ec2.model.Filter>)
func (jbobject *ServicesEc2ModelDescribeReservedInstancesOfferingsRequest) SetFilters(a []*ServicesEc2ModelFilter)  {
	conv_a := javabind.NewGoToJavaCollection(javabind.NewGoToJavaCallable())
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "setFilters", javabind.Void, conv_a.Value().Cast("java/util/Collection"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()

}

// public com.amazonaws.services.ec2.model.DescribeReservedInstancesOfferingsRequest withFilters(com.amazonaws.services.ec2.model.Filter...)
func (jbobject *ServicesEc2ModelDescribeReservedInstancesOfferingsRequest) WithFilters(a ...*ServicesEc2ModelFilter) *ServicesEc2ModelDescribeReservedInstancesOfferingsRequest {
	conv_a := javabind.NewGoToJavaObjectArray(javabind.NewGoToJavaCallable(), "com/amazonaws/services/ec2/model/Filter")
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "withFilters", "com/amazonaws/services/ec2/model/DescribeReservedInstancesOfferingsRequest", conv_a.Value().Cast("com/amazonaws/services/ec2/model/Filter"))
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
	unique_x := &ServicesEc2ModelDescribeReservedInstancesOfferingsRequest{}
	unique_x.Callable = dst
	return unique_x
}

// public com.amazonaws.services.ec2.model.DescribeReservedInstancesOfferingsRequest withFilters(java.util.Collection<com.amazonaws.services.ec2.model.Filter>)
func (jbobject *ServicesEc2ModelDescribeReservedInstancesOfferingsRequest) WithFilters2(a []*ServicesEc2ModelFilter) *ServicesEc2ModelDescribeReservedInstancesOfferingsRequest {
	conv_a := javabind.NewGoToJavaCollection(javabind.NewGoToJavaCallable())
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "withFilters", "com/amazonaws/services/ec2/model/DescribeReservedInstancesOfferingsRequest", conv_a.Value().Cast("java/util/Collection"))
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
	unique_x := &ServicesEc2ModelDescribeReservedInstancesOfferingsRequest{}
	unique_x.Callable = dst
	return unique_x
}

// public void setInstanceTenancy(java.lang.String)
func (jbobject *ServicesEc2ModelDescribeReservedInstancesOfferingsRequest) SetInstanceTenancy2(a string)  {
	conv_a := javabind.NewGoToJavaString()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "setInstanceTenancy", javabind.Void, conv_a.Value().Cast("java/lang/String"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()

}

// public java.lang.String getInstanceTenancy()
func (jbobject *ServicesEc2ModelDescribeReservedInstancesOfferingsRequest) GetInstanceTenancy() string {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getInstanceTenancy", "java/lang/String")
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

// public com.amazonaws.services.ec2.model.DescribeReservedInstancesOfferingsRequest withInstanceTenancy(java.lang.String)
func (jbobject *ServicesEc2ModelDescribeReservedInstancesOfferingsRequest) WithInstanceTenancy2(a string) *ServicesEc2ModelDescribeReservedInstancesOfferingsRequest {
	conv_a := javabind.NewGoToJavaString()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "withInstanceTenancy", "com/amazonaws/services/ec2/model/DescribeReservedInstancesOfferingsRequest", conv_a.Value().Cast("java/lang/String"))
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
	unique_x := &ServicesEc2ModelDescribeReservedInstancesOfferingsRequest{}
	unique_x.Callable = dst
	return unique_x
}

// public void setInstanceTenancy(com.amazonaws.services.ec2.model.Tenancy)
func (jbobject *ServicesEc2ModelDescribeReservedInstancesOfferingsRequest) SetInstanceTenancy(a ServicesEc2ModelTenancyInterface)  {
	conv_a := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "setInstanceTenancy", javabind.Void, conv_a.Value().Cast("com/amazonaws/services/ec2/model/Tenancy"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()

}

// public com.amazonaws.services.ec2.model.DescribeReservedInstancesOfferingsRequest withInstanceTenancy(com.amazonaws.services.ec2.model.Tenancy)
func (jbobject *ServicesEc2ModelDescribeReservedInstancesOfferingsRequest) WithInstanceTenancy(a ServicesEc2ModelTenancyInterface) *ServicesEc2ModelDescribeReservedInstancesOfferingsRequest {
	conv_a := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "withInstanceTenancy", "com/amazonaws/services/ec2/model/DescribeReservedInstancesOfferingsRequest", conv_a.Value().Cast("com/amazonaws/services/ec2/model/Tenancy"))
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
	unique_x := &ServicesEc2ModelDescribeReservedInstancesOfferingsRequest{}
	unique_x.Callable = dst
	return unique_x
}

// public void setOfferingType(java.lang.String)
func (jbobject *ServicesEc2ModelDescribeReservedInstancesOfferingsRequest) SetOfferingType2(a string)  {
	conv_a := javabind.NewGoToJavaString()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "setOfferingType", javabind.Void, conv_a.Value().Cast("java/lang/String"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()

}

// public java.lang.String getOfferingType()
func (jbobject *ServicesEc2ModelDescribeReservedInstancesOfferingsRequest) GetOfferingType() string {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getOfferingType", "java/lang/String")
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

// public com.amazonaws.services.ec2.model.DescribeReservedInstancesOfferingsRequest withOfferingType(java.lang.String)
func (jbobject *ServicesEc2ModelDescribeReservedInstancesOfferingsRequest) WithOfferingType2(a string) *ServicesEc2ModelDescribeReservedInstancesOfferingsRequest {
	conv_a := javabind.NewGoToJavaString()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "withOfferingType", "com/amazonaws/services/ec2/model/DescribeReservedInstancesOfferingsRequest", conv_a.Value().Cast("java/lang/String"))
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
	unique_x := &ServicesEc2ModelDescribeReservedInstancesOfferingsRequest{}
	unique_x.Callable = dst
	return unique_x
}

// public void setOfferingType(com.amazonaws.services.ec2.model.OfferingTypeValues)
func (jbobject *ServicesEc2ModelDescribeReservedInstancesOfferingsRequest) SetOfferingType(a ServicesEc2ModelOfferingTypeValuesInterface)  {
	conv_a := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "setOfferingType", javabind.Void, conv_a.Value().Cast("com/amazonaws/services/ec2/model/OfferingTypeValues"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()

}

// public com.amazonaws.services.ec2.model.DescribeReservedInstancesOfferingsRequest withOfferingType(com.amazonaws.services.ec2.model.OfferingTypeValues)
func (jbobject *ServicesEc2ModelDescribeReservedInstancesOfferingsRequest) WithOfferingType(a ServicesEc2ModelOfferingTypeValuesInterface) *ServicesEc2ModelDescribeReservedInstancesOfferingsRequest {
	conv_a := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "withOfferingType", "com/amazonaws/services/ec2/model/DescribeReservedInstancesOfferingsRequest", conv_a.Value().Cast("com/amazonaws/services/ec2/model/OfferingTypeValues"))
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
	unique_x := &ServicesEc2ModelDescribeReservedInstancesOfferingsRequest{}
	unique_x.Callable = dst
	return unique_x
}

// public void setNextToken(java.lang.String)
func (jbobject *ServicesEc2ModelDescribeReservedInstancesOfferingsRequest) SetNextToken(a string)  {
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
func (jbobject *ServicesEc2ModelDescribeReservedInstancesOfferingsRequest) GetNextToken() string {
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

// public com.amazonaws.services.ec2.model.DescribeReservedInstancesOfferingsRequest withNextToken(java.lang.String)
func (jbobject *ServicesEc2ModelDescribeReservedInstancesOfferingsRequest) WithNextToken(a string) *ServicesEc2ModelDescribeReservedInstancesOfferingsRequest {
	conv_a := javabind.NewGoToJavaString()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "withNextToken", "com/amazonaws/services/ec2/model/DescribeReservedInstancesOfferingsRequest", conv_a.Value().Cast("java/lang/String"))
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
	unique_x := &ServicesEc2ModelDescribeReservedInstancesOfferingsRequest{}
	unique_x.Callable = dst
	return unique_x
}

// public void setMaxResults(java.lang.Integer)
func (jbobject *ServicesEc2ModelDescribeReservedInstancesOfferingsRequest) SetMaxResults(a int)  {
	conv_a := javabind.NewGoToJavaInteger()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "setMaxResults", javabind.Void, conv_a.Value().Cast("java/lang/Integer"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()

}

// public java.lang.Integer getMaxResults()
func (jbobject *ServicesEc2ModelDescribeReservedInstancesOfferingsRequest) GetMaxResults() int {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getMaxResults", "java/lang/Integer")
	if err != nil {
		panic(err)
	}
	retconv := javabind.NewJavaToGoInteger()
	dst := new(int)
	retconv.Dest(dst)
	if err := retconv.Convert(javabind.ObjectRef(jret)); err != nil {
		panic(err)
	}
	retconv.CleanUp()
	return *dst
}

// public com.amazonaws.services.ec2.model.DescribeReservedInstancesOfferingsRequest withMaxResults(java.lang.Integer)
func (jbobject *ServicesEc2ModelDescribeReservedInstancesOfferingsRequest) WithMaxResults(a int) *ServicesEc2ModelDescribeReservedInstancesOfferingsRequest {
	conv_a := javabind.NewGoToJavaInteger()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "withMaxResults", "com/amazonaws/services/ec2/model/DescribeReservedInstancesOfferingsRequest", conv_a.Value().Cast("java/lang/Integer"))
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
	unique_x := &ServicesEc2ModelDescribeReservedInstancesOfferingsRequest{}
	unique_x.Callable = dst
	return unique_x
}

// public void setIncludeMarketplace(java.lang.Boolean)
func (jbobject *ServicesEc2ModelDescribeReservedInstancesOfferingsRequest) SetIncludeMarketplace(a bool)  {
	conv_a := javabind.NewGoToJavaBoolean()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "setIncludeMarketplace", javabind.Void, conv_a.Value().Cast("java/lang/Boolean"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()

}

// public java.lang.Boolean getIncludeMarketplace()
func (jbobject *ServicesEc2ModelDescribeReservedInstancesOfferingsRequest) GetIncludeMarketplace() bool {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getIncludeMarketplace", "java/lang/Boolean")
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

// public com.amazonaws.services.ec2.model.DescribeReservedInstancesOfferingsRequest withIncludeMarketplace(java.lang.Boolean)
func (jbobject *ServicesEc2ModelDescribeReservedInstancesOfferingsRequest) WithIncludeMarketplace(a bool) *ServicesEc2ModelDescribeReservedInstancesOfferingsRequest {
	conv_a := javabind.NewGoToJavaBoolean()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "withIncludeMarketplace", "com/amazonaws/services/ec2/model/DescribeReservedInstancesOfferingsRequest", conv_a.Value().Cast("java/lang/Boolean"))
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
	unique_x := &ServicesEc2ModelDescribeReservedInstancesOfferingsRequest{}
	unique_x.Callable = dst
	return unique_x
}

// public java.lang.Boolean isIncludeMarketplace()
func (jbobject *ServicesEc2ModelDescribeReservedInstancesOfferingsRequest) IsIncludeMarketplace() bool {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "isIncludeMarketplace", "java/lang/Boolean")
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

// public void setMinDuration(java.lang.Long)
func (jbobject *ServicesEc2ModelDescribeReservedInstancesOfferingsRequest) SetMinDuration(a int64)  {
	conv_a := javabind.NewGoToJavaLong()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "setMinDuration", javabind.Void, conv_a.Value().Cast("java/lang/Long"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()

}

// public java.lang.Long getMinDuration()
func (jbobject *ServicesEc2ModelDescribeReservedInstancesOfferingsRequest) GetMinDuration() int64 {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getMinDuration", "java/lang/Long")
	if err != nil {
		panic(err)
	}
	retconv := javabind.NewJavaToGoLong()
	dst := new(int64)
	retconv.Dest(dst)
	if err := retconv.Convert(javabind.ObjectRef(jret)); err != nil {
		panic(err)
	}
	retconv.CleanUp()
	return *dst
}

// public com.amazonaws.services.ec2.model.DescribeReservedInstancesOfferingsRequest withMinDuration(java.lang.Long)
func (jbobject *ServicesEc2ModelDescribeReservedInstancesOfferingsRequest) WithMinDuration(a int64) *ServicesEc2ModelDescribeReservedInstancesOfferingsRequest {
	conv_a := javabind.NewGoToJavaLong()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "withMinDuration", "com/amazonaws/services/ec2/model/DescribeReservedInstancesOfferingsRequest", conv_a.Value().Cast("java/lang/Long"))
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
	unique_x := &ServicesEc2ModelDescribeReservedInstancesOfferingsRequest{}
	unique_x.Callable = dst
	return unique_x
}

// public void setMaxDuration(java.lang.Long)
func (jbobject *ServicesEc2ModelDescribeReservedInstancesOfferingsRequest) SetMaxDuration(a int64)  {
	conv_a := javabind.NewGoToJavaLong()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "setMaxDuration", javabind.Void, conv_a.Value().Cast("java/lang/Long"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()

}

// public java.lang.Long getMaxDuration()
func (jbobject *ServicesEc2ModelDescribeReservedInstancesOfferingsRequest) GetMaxDuration() int64 {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getMaxDuration", "java/lang/Long")
	if err != nil {
		panic(err)
	}
	retconv := javabind.NewJavaToGoLong()
	dst := new(int64)
	retconv.Dest(dst)
	if err := retconv.Convert(javabind.ObjectRef(jret)); err != nil {
		panic(err)
	}
	retconv.CleanUp()
	return *dst
}

// public com.amazonaws.services.ec2.model.DescribeReservedInstancesOfferingsRequest withMaxDuration(java.lang.Long)
func (jbobject *ServicesEc2ModelDescribeReservedInstancesOfferingsRequest) WithMaxDuration(a int64) *ServicesEc2ModelDescribeReservedInstancesOfferingsRequest {
	conv_a := javabind.NewGoToJavaLong()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "withMaxDuration", "com/amazonaws/services/ec2/model/DescribeReservedInstancesOfferingsRequest", conv_a.Value().Cast("java/lang/Long"))
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
	unique_x := &ServicesEc2ModelDescribeReservedInstancesOfferingsRequest{}
	unique_x.Callable = dst
	return unique_x
}

// public void setMaxInstanceCount(java.lang.Integer)
func (jbobject *ServicesEc2ModelDescribeReservedInstancesOfferingsRequest) SetMaxInstanceCount(a int)  {
	conv_a := javabind.NewGoToJavaInteger()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "setMaxInstanceCount", javabind.Void, conv_a.Value().Cast("java/lang/Integer"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()

}

// public java.lang.Integer getMaxInstanceCount()
func (jbobject *ServicesEc2ModelDescribeReservedInstancesOfferingsRequest) GetMaxInstanceCount() int {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getMaxInstanceCount", "java/lang/Integer")
	if err != nil {
		panic(err)
	}
	retconv := javabind.NewJavaToGoInteger()
	dst := new(int)
	retconv.Dest(dst)
	if err := retconv.Convert(javabind.ObjectRef(jret)); err != nil {
		panic(err)
	}
	retconv.CleanUp()
	return *dst
}

// public com.amazonaws.services.ec2.model.DescribeReservedInstancesOfferingsRequest withMaxInstanceCount(java.lang.Integer)
func (jbobject *ServicesEc2ModelDescribeReservedInstancesOfferingsRequest) WithMaxInstanceCount(a int) *ServicesEc2ModelDescribeReservedInstancesOfferingsRequest {
	conv_a := javabind.NewGoToJavaInteger()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "withMaxInstanceCount", "com/amazonaws/services/ec2/model/DescribeReservedInstancesOfferingsRequest", conv_a.Value().Cast("java/lang/Integer"))
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
	unique_x := &ServicesEc2ModelDescribeReservedInstancesOfferingsRequest{}
	unique_x.Callable = dst
	return unique_x
}

// public com.amazonaws.Request<com.amazonaws.services.ec2.model.DescribeReservedInstancesOfferingsRequest> getDryRunRequest()
func (jbobject *ServicesEc2ModelDescribeReservedInstancesOfferingsRequest) GetDryRunRequest() *Request {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getDryRunRequest", "com/amazonaws/Request")
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
	unique_x := &Request{}
	unique_x.Callable = dst
	return unique_x
}

// public java.lang.String toString()
func (jbobject *ServicesEc2ModelDescribeReservedInstancesOfferingsRequest) ToString() string {
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
func (jbobject *ServicesEc2ModelDescribeReservedInstancesOfferingsRequest) Equals(a interface{}) bool {
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
func (jbobject *ServicesEc2ModelDescribeReservedInstancesOfferingsRequest) HashCode() int {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "hashCode", javabind.Int)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

// public com.amazonaws.services.ec2.model.DescribeReservedInstancesOfferingsRequest clone()
func (jbobject *ServicesEc2ModelDescribeReservedInstancesOfferingsRequest) Clone3() *ServicesEc2ModelDescribeReservedInstancesOfferingsRequest {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "clone", "com/amazonaws/services/ec2/model/DescribeReservedInstancesOfferingsRequest")
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
	unique_x := &ServicesEc2ModelDescribeReservedInstancesOfferingsRequest{}
	unique_x.Callable = dst
	return unique_x
}

// public com.amazonaws.AmazonWebServiceRequest clone()
func (jbobject *ServicesEc2ModelDescribeReservedInstancesOfferingsRequest) Clone() *AmazonWebServiceRequest {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "clone", "com/amazonaws/AmazonWebServiceRequest")
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
	unique_x := &AmazonWebServiceRequest{}
	unique_x.Callable = dst
	return unique_x
}

// public java.lang.Object clone() throws java.lang.CloneNotSupportedException
func (jbobject *ServicesEc2ModelDescribeReservedInstancesOfferingsRequest) Clone2() (*JavaLangObject, error) {
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


