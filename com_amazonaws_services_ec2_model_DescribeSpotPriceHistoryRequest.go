package ec2

import "time"
import "github.com/timob/javabind"

type ServicesEc2ModelDescribeSpotPriceHistoryRequestInterface interface {
	AmazonWebServiceRequestInterface

	// public void setStartTime(java.util.Date)
	SetStartTime(a time.Time) 

	// public java.util.Date getStartTime()
	GetStartTime() time.Time

	// public com.amazonaws.services.ec2.model.DescribeSpotPriceHistoryRequest withStartTime(java.util.Date)
	WithStartTime(a time.Time) *ServicesEc2ModelDescribeSpotPriceHistoryRequest

	// public void setEndTime(java.util.Date)
	SetEndTime(a time.Time) 

	// public java.util.Date getEndTime()
	GetEndTime() time.Time

	// public com.amazonaws.services.ec2.model.DescribeSpotPriceHistoryRequest withEndTime(java.util.Date)
	WithEndTime(a time.Time) *ServicesEc2ModelDescribeSpotPriceHistoryRequest

	// public java.util.List<java.lang.String> getInstanceTypes()
	GetInstanceTypes() []string

	// public void setInstanceTypes(java.util.Collection<java.lang.String>)
	SetInstanceTypes(a []string) 

	// public com.amazonaws.services.ec2.model.DescribeSpotPriceHistoryRequest withInstanceTypes(java.lang.String...)
	WithInstanceTypes2(a ...string) *ServicesEc2ModelDescribeSpotPriceHistoryRequest

	// public com.amazonaws.services.ec2.model.DescribeSpotPriceHistoryRequest withInstanceTypes(java.util.Collection<java.lang.String>)
	WithInstanceTypes3(a []string) *ServicesEc2ModelDescribeSpotPriceHistoryRequest

	// public com.amazonaws.services.ec2.model.DescribeSpotPriceHistoryRequest withInstanceTypes(com.amazonaws.services.ec2.model.InstanceType...)
	WithInstanceTypes(a ...*ServicesEc2ModelInstanceType) *ServicesEc2ModelDescribeSpotPriceHistoryRequest

	// public java.util.List<java.lang.String> getProductDescriptions()
	GetProductDescriptions() []string

	// public void setProductDescriptions(java.util.Collection<java.lang.String>)
	SetProductDescriptions(a []string) 

	// public com.amazonaws.services.ec2.model.DescribeSpotPriceHistoryRequest withProductDescriptions(java.lang.String...)
	WithProductDescriptions(a ...string) *ServicesEc2ModelDescribeSpotPriceHistoryRequest

	// public com.amazonaws.services.ec2.model.DescribeSpotPriceHistoryRequest withProductDescriptions(java.util.Collection<java.lang.String>)
	WithProductDescriptions2(a []string) *ServicesEc2ModelDescribeSpotPriceHistoryRequest

	// public java.util.List<com.amazonaws.services.ec2.model.Filter> getFilters()
	GetFilters() []*ServicesEc2ModelFilter

	// public void setFilters(java.util.Collection<com.amazonaws.services.ec2.model.Filter>)
	SetFilters(a []*ServicesEc2ModelFilter) 

	// public com.amazonaws.services.ec2.model.DescribeSpotPriceHistoryRequest withFilters(com.amazonaws.services.ec2.model.Filter...)
	WithFilters(a ...*ServicesEc2ModelFilter) *ServicesEc2ModelDescribeSpotPriceHistoryRequest

	// public com.amazonaws.services.ec2.model.DescribeSpotPriceHistoryRequest withFilters(java.util.Collection<com.amazonaws.services.ec2.model.Filter>)
	WithFilters2(a []*ServicesEc2ModelFilter) *ServicesEc2ModelDescribeSpotPriceHistoryRequest

	// public void setAvailabilityZone(java.lang.String)
	SetAvailabilityZone(a string) 

	// public java.lang.String getAvailabilityZone()
	GetAvailabilityZone() string

	// public com.amazonaws.services.ec2.model.DescribeSpotPriceHistoryRequest withAvailabilityZone(java.lang.String)
	WithAvailabilityZone(a string) *ServicesEc2ModelDescribeSpotPriceHistoryRequest

	// public void setMaxResults(java.lang.Integer)
	SetMaxResults(a int) 

	// public java.lang.Integer getMaxResults()
	GetMaxResults() int

	// public com.amazonaws.services.ec2.model.DescribeSpotPriceHistoryRequest withMaxResults(java.lang.Integer)
	WithMaxResults(a int) *ServicesEc2ModelDescribeSpotPriceHistoryRequest

	// public void setNextToken(java.lang.String)
	SetNextToken(a string) 

	// public java.lang.String getNextToken()
	GetNextToken() string

	// public com.amazonaws.services.ec2.model.DescribeSpotPriceHistoryRequest withNextToken(java.lang.String)
	WithNextToken(a string) *ServicesEc2ModelDescribeSpotPriceHistoryRequest

	// public com.amazonaws.Request<com.amazonaws.services.ec2.model.DescribeSpotPriceHistoryRequest> getDryRunRequest()
	GetDryRunRequest() *Request

	// public com.amazonaws.services.ec2.model.DescribeSpotPriceHistoryRequest clone()
	Clone3() *ServicesEc2ModelDescribeSpotPriceHistoryRequest
}

type ServicesEc2ModelDescribeSpotPriceHistoryRequest struct {
	AmazonWebServiceRequest
}

// public com.amazonaws.services.ec2.model.DescribeSpotPriceHistoryRequest()
func NewServicesEc2ModelDescribeSpotPriceHistoryRequest() (*ServicesEc2ModelDescribeSpotPriceHistoryRequest) {

	obj, err := javabind.GetEnv().NewObject("com/amazonaws/services/ec2/model/DescribeSpotPriceHistoryRequest")
	if err != nil {
		panic(err)
	}
	x := &ServicesEc2ModelDescribeSpotPriceHistoryRequest{}
	x.Callable = &javabind.Callable{obj}
	return x
}

// public void setStartTime(java.util.Date)
func (jbobject *ServicesEc2ModelDescribeSpotPriceHistoryRequest) SetStartTime(a time.Time)  {
	conv_a := javabind.NewGoToJavaDate()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "setStartTime", javabind.Void, conv_a.Value().Cast("java/util/Date"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()

}

// public java.util.Date getStartTime()
func (jbobject *ServicesEc2ModelDescribeSpotPriceHistoryRequest) GetStartTime() time.Time {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getStartTime", "java/util/Date")
	if err != nil {
		panic(err)
	}
	retconv := javabind.NewJavaToGoDate()
	dst := new(time.Time)
	retconv.Dest(dst)
	if err := retconv.Convert(javabind.ObjectRef(jret)); err != nil {
		panic(err)
	}
	retconv.CleanUp()
	return *dst
}

// public com.amazonaws.services.ec2.model.DescribeSpotPriceHistoryRequest withStartTime(java.util.Date)
func (jbobject *ServicesEc2ModelDescribeSpotPriceHistoryRequest) WithStartTime(a time.Time) *ServicesEc2ModelDescribeSpotPriceHistoryRequest {
	conv_a := javabind.NewGoToJavaDate()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "withStartTime", "com/amazonaws/services/ec2/model/DescribeSpotPriceHistoryRequest", conv_a.Value().Cast("java/util/Date"))
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
	unique_x := &ServicesEc2ModelDescribeSpotPriceHistoryRequest{}
	unique_x.Callable = dst
	return unique_x
}

// public void setEndTime(java.util.Date)
func (jbobject *ServicesEc2ModelDescribeSpotPriceHistoryRequest) SetEndTime(a time.Time)  {
	conv_a := javabind.NewGoToJavaDate()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "setEndTime", javabind.Void, conv_a.Value().Cast("java/util/Date"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()

}

// public java.util.Date getEndTime()
func (jbobject *ServicesEc2ModelDescribeSpotPriceHistoryRequest) GetEndTime() time.Time {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getEndTime", "java/util/Date")
	if err != nil {
		panic(err)
	}
	retconv := javabind.NewJavaToGoDate()
	dst := new(time.Time)
	retconv.Dest(dst)
	if err := retconv.Convert(javabind.ObjectRef(jret)); err != nil {
		panic(err)
	}
	retconv.CleanUp()
	return *dst
}

// public com.amazonaws.services.ec2.model.DescribeSpotPriceHistoryRequest withEndTime(java.util.Date)
func (jbobject *ServicesEc2ModelDescribeSpotPriceHistoryRequest) WithEndTime(a time.Time) *ServicesEc2ModelDescribeSpotPriceHistoryRequest {
	conv_a := javabind.NewGoToJavaDate()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "withEndTime", "com/amazonaws/services/ec2/model/DescribeSpotPriceHistoryRequest", conv_a.Value().Cast("java/util/Date"))
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
	unique_x := &ServicesEc2ModelDescribeSpotPriceHistoryRequest{}
	unique_x.Callable = dst
	return unique_x
}

// public java.util.List<java.lang.String> getInstanceTypes()
func (jbobject *ServicesEc2ModelDescribeSpotPriceHistoryRequest) GetInstanceTypes() []string {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getInstanceTypes", "java/util/List")
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

// public void setInstanceTypes(java.util.Collection<java.lang.String>)
func (jbobject *ServicesEc2ModelDescribeSpotPriceHistoryRequest) SetInstanceTypes(a []string)  {
	conv_a := javabind.NewGoToJavaCollection(javabind.NewGoToJavaString())
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "setInstanceTypes", javabind.Void, conv_a.Value().Cast("java/util/Collection"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()

}

// public com.amazonaws.services.ec2.model.DescribeSpotPriceHistoryRequest withInstanceTypes(java.lang.String...)
func (jbobject *ServicesEc2ModelDescribeSpotPriceHistoryRequest) WithInstanceTypes2(a ...string) *ServicesEc2ModelDescribeSpotPriceHistoryRequest {
	conv_a := javabind.NewGoToJavaObjectArray(javabind.NewGoToJavaString(), "java/lang/String")
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "withInstanceTypes", "com/amazonaws/services/ec2/model/DescribeSpotPriceHistoryRequest", conv_a.Value().Cast("java/lang/String"))
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
	unique_x := &ServicesEc2ModelDescribeSpotPriceHistoryRequest{}
	unique_x.Callable = dst
	return unique_x
}

// public com.amazonaws.services.ec2.model.DescribeSpotPriceHistoryRequest withInstanceTypes(java.util.Collection<java.lang.String>)
func (jbobject *ServicesEc2ModelDescribeSpotPriceHistoryRequest) WithInstanceTypes3(a []string) *ServicesEc2ModelDescribeSpotPriceHistoryRequest {
	conv_a := javabind.NewGoToJavaCollection(javabind.NewGoToJavaString())
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "withInstanceTypes", "com/amazonaws/services/ec2/model/DescribeSpotPriceHistoryRequest", conv_a.Value().Cast("java/util/Collection"))
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
	unique_x := &ServicesEc2ModelDescribeSpotPriceHistoryRequest{}
	unique_x.Callable = dst
	return unique_x
}

// public com.amazonaws.services.ec2.model.DescribeSpotPriceHistoryRequest withInstanceTypes(com.amazonaws.services.ec2.model.InstanceType...)
func (jbobject *ServicesEc2ModelDescribeSpotPriceHistoryRequest) WithInstanceTypes(a ...*ServicesEc2ModelInstanceType) *ServicesEc2ModelDescribeSpotPriceHistoryRequest {
	conv_a := javabind.NewGoToJavaObjectArray(javabind.NewGoToJavaCallable(), "com/amazonaws/services/ec2/model/InstanceType")
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "withInstanceTypes", "com/amazonaws/services/ec2/model/DescribeSpotPriceHistoryRequest", conv_a.Value().Cast("com/amazonaws/services/ec2/model/InstanceType"))
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
	unique_x := &ServicesEc2ModelDescribeSpotPriceHistoryRequest{}
	unique_x.Callable = dst
	return unique_x
}

// public java.util.List<java.lang.String> getProductDescriptions()
func (jbobject *ServicesEc2ModelDescribeSpotPriceHistoryRequest) GetProductDescriptions() []string {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getProductDescriptions", "java/util/List")
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

// public void setProductDescriptions(java.util.Collection<java.lang.String>)
func (jbobject *ServicesEc2ModelDescribeSpotPriceHistoryRequest) SetProductDescriptions(a []string)  {
	conv_a := javabind.NewGoToJavaCollection(javabind.NewGoToJavaString())
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "setProductDescriptions", javabind.Void, conv_a.Value().Cast("java/util/Collection"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()

}

// public com.amazonaws.services.ec2.model.DescribeSpotPriceHistoryRequest withProductDescriptions(java.lang.String...)
func (jbobject *ServicesEc2ModelDescribeSpotPriceHistoryRequest) WithProductDescriptions(a ...string) *ServicesEc2ModelDescribeSpotPriceHistoryRequest {
	conv_a := javabind.NewGoToJavaObjectArray(javabind.NewGoToJavaString(), "java/lang/String")
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "withProductDescriptions", "com/amazonaws/services/ec2/model/DescribeSpotPriceHistoryRequest", conv_a.Value().Cast("java/lang/String"))
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
	unique_x := &ServicesEc2ModelDescribeSpotPriceHistoryRequest{}
	unique_x.Callable = dst
	return unique_x
}

// public com.amazonaws.services.ec2.model.DescribeSpotPriceHistoryRequest withProductDescriptions(java.util.Collection<java.lang.String>)
func (jbobject *ServicesEc2ModelDescribeSpotPriceHistoryRequest) WithProductDescriptions2(a []string) *ServicesEc2ModelDescribeSpotPriceHistoryRequest {
	conv_a := javabind.NewGoToJavaCollection(javabind.NewGoToJavaString())
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "withProductDescriptions", "com/amazonaws/services/ec2/model/DescribeSpotPriceHistoryRequest", conv_a.Value().Cast("java/util/Collection"))
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
	unique_x := &ServicesEc2ModelDescribeSpotPriceHistoryRequest{}
	unique_x.Callable = dst
	return unique_x
}

// public java.util.List<com.amazonaws.services.ec2.model.Filter> getFilters()
func (jbobject *ServicesEc2ModelDescribeSpotPriceHistoryRequest) GetFilters() []*ServicesEc2ModelFilter {
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
func (jbobject *ServicesEc2ModelDescribeSpotPriceHistoryRequest) SetFilters(a []*ServicesEc2ModelFilter)  {
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

// public com.amazonaws.services.ec2.model.DescribeSpotPriceHistoryRequest withFilters(com.amazonaws.services.ec2.model.Filter...)
func (jbobject *ServicesEc2ModelDescribeSpotPriceHistoryRequest) WithFilters(a ...*ServicesEc2ModelFilter) *ServicesEc2ModelDescribeSpotPriceHistoryRequest {
	conv_a := javabind.NewGoToJavaObjectArray(javabind.NewGoToJavaCallable(), "com/amazonaws/services/ec2/model/Filter")
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "withFilters", "com/amazonaws/services/ec2/model/DescribeSpotPriceHistoryRequest", conv_a.Value().Cast("com/amazonaws/services/ec2/model/Filter"))
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
	unique_x := &ServicesEc2ModelDescribeSpotPriceHistoryRequest{}
	unique_x.Callable = dst
	return unique_x
}

// public com.amazonaws.services.ec2.model.DescribeSpotPriceHistoryRequest withFilters(java.util.Collection<com.amazonaws.services.ec2.model.Filter>)
func (jbobject *ServicesEc2ModelDescribeSpotPriceHistoryRequest) WithFilters2(a []*ServicesEc2ModelFilter) *ServicesEc2ModelDescribeSpotPriceHistoryRequest {
	conv_a := javabind.NewGoToJavaCollection(javabind.NewGoToJavaCallable())
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "withFilters", "com/amazonaws/services/ec2/model/DescribeSpotPriceHistoryRequest", conv_a.Value().Cast("java/util/Collection"))
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
	unique_x := &ServicesEc2ModelDescribeSpotPriceHistoryRequest{}
	unique_x.Callable = dst
	return unique_x
}

// public void setAvailabilityZone(java.lang.String)
func (jbobject *ServicesEc2ModelDescribeSpotPriceHistoryRequest) SetAvailabilityZone(a string)  {
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
func (jbobject *ServicesEc2ModelDescribeSpotPriceHistoryRequest) GetAvailabilityZone() string {
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

// public com.amazonaws.services.ec2.model.DescribeSpotPriceHistoryRequest withAvailabilityZone(java.lang.String)
func (jbobject *ServicesEc2ModelDescribeSpotPriceHistoryRequest) WithAvailabilityZone(a string) *ServicesEc2ModelDescribeSpotPriceHistoryRequest {
	conv_a := javabind.NewGoToJavaString()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "withAvailabilityZone", "com/amazonaws/services/ec2/model/DescribeSpotPriceHistoryRequest", conv_a.Value().Cast("java/lang/String"))
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
	unique_x := &ServicesEc2ModelDescribeSpotPriceHistoryRequest{}
	unique_x.Callable = dst
	return unique_x
}

// public void setMaxResults(java.lang.Integer)
func (jbobject *ServicesEc2ModelDescribeSpotPriceHistoryRequest) SetMaxResults(a int)  {
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
func (jbobject *ServicesEc2ModelDescribeSpotPriceHistoryRequest) GetMaxResults() int {
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

// public com.amazonaws.services.ec2.model.DescribeSpotPriceHistoryRequest withMaxResults(java.lang.Integer)
func (jbobject *ServicesEc2ModelDescribeSpotPriceHistoryRequest) WithMaxResults(a int) *ServicesEc2ModelDescribeSpotPriceHistoryRequest {
	conv_a := javabind.NewGoToJavaInteger()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "withMaxResults", "com/amazonaws/services/ec2/model/DescribeSpotPriceHistoryRequest", conv_a.Value().Cast("java/lang/Integer"))
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
	unique_x := &ServicesEc2ModelDescribeSpotPriceHistoryRequest{}
	unique_x.Callable = dst
	return unique_x
}

// public void setNextToken(java.lang.String)
func (jbobject *ServicesEc2ModelDescribeSpotPriceHistoryRequest) SetNextToken(a string)  {
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
func (jbobject *ServicesEc2ModelDescribeSpotPriceHistoryRequest) GetNextToken() string {
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

// public com.amazonaws.services.ec2.model.DescribeSpotPriceHistoryRequest withNextToken(java.lang.String)
func (jbobject *ServicesEc2ModelDescribeSpotPriceHistoryRequest) WithNextToken(a string) *ServicesEc2ModelDescribeSpotPriceHistoryRequest {
	conv_a := javabind.NewGoToJavaString()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "withNextToken", "com/amazonaws/services/ec2/model/DescribeSpotPriceHistoryRequest", conv_a.Value().Cast("java/lang/String"))
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
	unique_x := &ServicesEc2ModelDescribeSpotPriceHistoryRequest{}
	unique_x.Callable = dst
	return unique_x
}

// public com.amazonaws.Request<com.amazonaws.services.ec2.model.DescribeSpotPriceHistoryRequest> getDryRunRequest()
func (jbobject *ServicesEc2ModelDescribeSpotPriceHistoryRequest) GetDryRunRequest() *Request {
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
func (jbobject *ServicesEc2ModelDescribeSpotPriceHistoryRequest) ToString() string {
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
func (jbobject *ServicesEc2ModelDescribeSpotPriceHistoryRequest) Equals(a interface{}) bool {
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
func (jbobject *ServicesEc2ModelDescribeSpotPriceHistoryRequest) HashCode() int {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "hashCode", javabind.Int)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

// public com.amazonaws.services.ec2.model.DescribeSpotPriceHistoryRequest clone()
func (jbobject *ServicesEc2ModelDescribeSpotPriceHistoryRequest) Clone3() *ServicesEc2ModelDescribeSpotPriceHistoryRequest {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "clone", "com/amazonaws/services/ec2/model/DescribeSpotPriceHistoryRequest")
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
	unique_x := &ServicesEc2ModelDescribeSpotPriceHistoryRequest{}
	unique_x.Callable = dst
	return unique_x
}

// public com.amazonaws.AmazonWebServiceRequest clone()
func (jbobject *ServicesEc2ModelDescribeSpotPriceHistoryRequest) Clone() *AmazonWebServiceRequest {
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
func (jbobject *ServicesEc2ModelDescribeSpotPriceHistoryRequest) Clone2() (*JavaLangObject, error) {
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


