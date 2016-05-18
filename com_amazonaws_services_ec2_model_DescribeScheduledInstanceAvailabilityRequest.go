package ec2

import "github.com/timob/javabind"

type ServicesEc2ModelDescribeScheduledInstanceAvailabilityRequestInterface interface {
	AmazonWebServiceRequestInterface

	// public void setRecurrence(com.amazonaws.services.ec2.model.ScheduledInstanceRecurrenceRequest)
	SetRecurrence(a ServicesEc2ModelScheduledInstanceRecurrenceRequestInterface) 

	// public com.amazonaws.services.ec2.model.ScheduledInstanceRecurrenceRequest getRecurrence()
	GetRecurrence() *ServicesEc2ModelScheduledInstanceRecurrenceRequest

	// public com.amazonaws.services.ec2.model.DescribeScheduledInstanceAvailabilityRequest withRecurrence(com.amazonaws.services.ec2.model.ScheduledInstanceRecurrenceRequest)
	WithRecurrence(a ServicesEc2ModelScheduledInstanceRecurrenceRequestInterface) *ServicesEc2ModelDescribeScheduledInstanceAvailabilityRequest

	// public void setFirstSlotStartTimeRange(com.amazonaws.services.ec2.model.SlotDateTimeRangeRequest)
	SetFirstSlotStartTimeRange(a ServicesEc2ModelSlotDateTimeRangeRequestInterface) 

	// public com.amazonaws.services.ec2.model.SlotDateTimeRangeRequest getFirstSlotStartTimeRange()
	GetFirstSlotStartTimeRange() *ServicesEc2ModelSlotDateTimeRangeRequest

	// public com.amazonaws.services.ec2.model.DescribeScheduledInstanceAvailabilityRequest withFirstSlotStartTimeRange(com.amazonaws.services.ec2.model.SlotDateTimeRangeRequest)
	WithFirstSlotStartTimeRange(a ServicesEc2ModelSlotDateTimeRangeRequestInterface) *ServicesEc2ModelDescribeScheduledInstanceAvailabilityRequest

	// public void setMinSlotDurationInHours(java.lang.Integer)
	SetMinSlotDurationInHours(a int) 

	// public java.lang.Integer getMinSlotDurationInHours()
	GetMinSlotDurationInHours() int

	// public com.amazonaws.services.ec2.model.DescribeScheduledInstanceAvailabilityRequest withMinSlotDurationInHours(java.lang.Integer)
	WithMinSlotDurationInHours(a int) *ServicesEc2ModelDescribeScheduledInstanceAvailabilityRequest

	// public void setMaxSlotDurationInHours(java.lang.Integer)
	SetMaxSlotDurationInHours(a int) 

	// public java.lang.Integer getMaxSlotDurationInHours()
	GetMaxSlotDurationInHours() int

	// public com.amazonaws.services.ec2.model.DescribeScheduledInstanceAvailabilityRequest withMaxSlotDurationInHours(java.lang.Integer)
	WithMaxSlotDurationInHours(a int) *ServicesEc2ModelDescribeScheduledInstanceAvailabilityRequest

	// public void setNextToken(java.lang.String)
	SetNextToken(a string) 

	// public java.lang.String getNextToken()
	GetNextToken() string

	// public com.amazonaws.services.ec2.model.DescribeScheduledInstanceAvailabilityRequest withNextToken(java.lang.String)
	WithNextToken(a string) *ServicesEc2ModelDescribeScheduledInstanceAvailabilityRequest

	// public void setMaxResults(java.lang.Integer)
	SetMaxResults(a int) 

	// public java.lang.Integer getMaxResults()
	GetMaxResults() int

	// public com.amazonaws.services.ec2.model.DescribeScheduledInstanceAvailabilityRequest withMaxResults(java.lang.Integer)
	WithMaxResults(a int) *ServicesEc2ModelDescribeScheduledInstanceAvailabilityRequest

	// public java.util.List<com.amazonaws.services.ec2.model.Filter> getFilters()
	GetFilters() []*ServicesEc2ModelFilter

	// public void setFilters(java.util.Collection<com.amazonaws.services.ec2.model.Filter>)
	SetFilters(a []*ServicesEc2ModelFilter) 

	// public com.amazonaws.services.ec2.model.DescribeScheduledInstanceAvailabilityRequest withFilters(com.amazonaws.services.ec2.model.Filter...)
	WithFilters(a ...*ServicesEc2ModelFilter) *ServicesEc2ModelDescribeScheduledInstanceAvailabilityRequest

	// public com.amazonaws.services.ec2.model.DescribeScheduledInstanceAvailabilityRequest withFilters(java.util.Collection<com.amazonaws.services.ec2.model.Filter>)
	WithFilters2(a []*ServicesEc2ModelFilter) *ServicesEc2ModelDescribeScheduledInstanceAvailabilityRequest

	// public com.amazonaws.Request<com.amazonaws.services.ec2.model.DescribeScheduledInstanceAvailabilityRequest> getDryRunRequest()
	GetDryRunRequest() *Request

	// public com.amazonaws.services.ec2.model.DescribeScheduledInstanceAvailabilityRequest clone()
	Clone3() *ServicesEc2ModelDescribeScheduledInstanceAvailabilityRequest
}

type ServicesEc2ModelDescribeScheduledInstanceAvailabilityRequest struct {
	AmazonWebServiceRequest
}

// public com.amazonaws.services.ec2.model.DescribeScheduledInstanceAvailabilityRequest()
func NewServicesEc2ModelDescribeScheduledInstanceAvailabilityRequest() (*ServicesEc2ModelDescribeScheduledInstanceAvailabilityRequest) {

	obj, err := javabind.GetEnv().NewObject("com/amazonaws/services/ec2/model/DescribeScheduledInstanceAvailabilityRequest")
	if err != nil {
		panic(err)
	}
	x := &ServicesEc2ModelDescribeScheduledInstanceAvailabilityRequest{}
	x.Callable = &javabind.Callable{obj}
	return x
}

// public void setRecurrence(com.amazonaws.services.ec2.model.ScheduledInstanceRecurrenceRequest)
func (jbobject *ServicesEc2ModelDescribeScheduledInstanceAvailabilityRequest) SetRecurrence(a ServicesEc2ModelScheduledInstanceRecurrenceRequestInterface)  {
	conv_a := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "setRecurrence", javabind.Void, conv_a.Value().Cast("com/amazonaws/services/ec2/model/ScheduledInstanceRecurrenceRequest"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()

}

// public com.amazonaws.services.ec2.model.ScheduledInstanceRecurrenceRequest getRecurrence()
func (jbobject *ServicesEc2ModelDescribeScheduledInstanceAvailabilityRequest) GetRecurrence() *ServicesEc2ModelScheduledInstanceRecurrenceRequest {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getRecurrence", "com/amazonaws/services/ec2/model/ScheduledInstanceRecurrenceRequest")
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
	unique_x := &ServicesEc2ModelScheduledInstanceRecurrenceRequest{}
	unique_x.Callable = dst
	return unique_x
}

// public com.amazonaws.services.ec2.model.DescribeScheduledInstanceAvailabilityRequest withRecurrence(com.amazonaws.services.ec2.model.ScheduledInstanceRecurrenceRequest)
func (jbobject *ServicesEc2ModelDescribeScheduledInstanceAvailabilityRequest) WithRecurrence(a ServicesEc2ModelScheduledInstanceRecurrenceRequestInterface) *ServicesEc2ModelDescribeScheduledInstanceAvailabilityRequest {
	conv_a := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "withRecurrence", "com/amazonaws/services/ec2/model/DescribeScheduledInstanceAvailabilityRequest", conv_a.Value().Cast("com/amazonaws/services/ec2/model/ScheduledInstanceRecurrenceRequest"))
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
	unique_x := &ServicesEc2ModelDescribeScheduledInstanceAvailabilityRequest{}
	unique_x.Callable = dst
	return unique_x
}

// public void setFirstSlotStartTimeRange(com.amazonaws.services.ec2.model.SlotDateTimeRangeRequest)
func (jbobject *ServicesEc2ModelDescribeScheduledInstanceAvailabilityRequest) SetFirstSlotStartTimeRange(a ServicesEc2ModelSlotDateTimeRangeRequestInterface)  {
	conv_a := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "setFirstSlotStartTimeRange", javabind.Void, conv_a.Value().Cast("com/amazonaws/services/ec2/model/SlotDateTimeRangeRequest"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()

}

// public com.amazonaws.services.ec2.model.SlotDateTimeRangeRequest getFirstSlotStartTimeRange()
func (jbobject *ServicesEc2ModelDescribeScheduledInstanceAvailabilityRequest) GetFirstSlotStartTimeRange() *ServicesEc2ModelSlotDateTimeRangeRequest {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getFirstSlotStartTimeRange", "com/amazonaws/services/ec2/model/SlotDateTimeRangeRequest")
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
	unique_x := &ServicesEc2ModelSlotDateTimeRangeRequest{}
	unique_x.Callable = dst
	return unique_x
}

// public com.amazonaws.services.ec2.model.DescribeScheduledInstanceAvailabilityRequest withFirstSlotStartTimeRange(com.amazonaws.services.ec2.model.SlotDateTimeRangeRequest)
func (jbobject *ServicesEc2ModelDescribeScheduledInstanceAvailabilityRequest) WithFirstSlotStartTimeRange(a ServicesEc2ModelSlotDateTimeRangeRequestInterface) *ServicesEc2ModelDescribeScheduledInstanceAvailabilityRequest {
	conv_a := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "withFirstSlotStartTimeRange", "com/amazonaws/services/ec2/model/DescribeScheduledInstanceAvailabilityRequest", conv_a.Value().Cast("com/amazonaws/services/ec2/model/SlotDateTimeRangeRequest"))
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
	unique_x := &ServicesEc2ModelDescribeScheduledInstanceAvailabilityRequest{}
	unique_x.Callable = dst
	return unique_x
}

// public void setMinSlotDurationInHours(java.lang.Integer)
func (jbobject *ServicesEc2ModelDescribeScheduledInstanceAvailabilityRequest) SetMinSlotDurationInHours(a int)  {
	conv_a := javabind.NewGoToJavaInteger()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "setMinSlotDurationInHours", javabind.Void, conv_a.Value().Cast("java/lang/Integer"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()

}

// public java.lang.Integer getMinSlotDurationInHours()
func (jbobject *ServicesEc2ModelDescribeScheduledInstanceAvailabilityRequest) GetMinSlotDurationInHours() int {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getMinSlotDurationInHours", "java/lang/Integer")
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

// public com.amazonaws.services.ec2.model.DescribeScheduledInstanceAvailabilityRequest withMinSlotDurationInHours(java.lang.Integer)
func (jbobject *ServicesEc2ModelDescribeScheduledInstanceAvailabilityRequest) WithMinSlotDurationInHours(a int) *ServicesEc2ModelDescribeScheduledInstanceAvailabilityRequest {
	conv_a := javabind.NewGoToJavaInteger()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "withMinSlotDurationInHours", "com/amazonaws/services/ec2/model/DescribeScheduledInstanceAvailabilityRequest", conv_a.Value().Cast("java/lang/Integer"))
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
	unique_x := &ServicesEc2ModelDescribeScheduledInstanceAvailabilityRequest{}
	unique_x.Callable = dst
	return unique_x
}

// public void setMaxSlotDurationInHours(java.lang.Integer)
func (jbobject *ServicesEc2ModelDescribeScheduledInstanceAvailabilityRequest) SetMaxSlotDurationInHours(a int)  {
	conv_a := javabind.NewGoToJavaInteger()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "setMaxSlotDurationInHours", javabind.Void, conv_a.Value().Cast("java/lang/Integer"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()

}

// public java.lang.Integer getMaxSlotDurationInHours()
func (jbobject *ServicesEc2ModelDescribeScheduledInstanceAvailabilityRequest) GetMaxSlotDurationInHours() int {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getMaxSlotDurationInHours", "java/lang/Integer")
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

// public com.amazonaws.services.ec2.model.DescribeScheduledInstanceAvailabilityRequest withMaxSlotDurationInHours(java.lang.Integer)
func (jbobject *ServicesEc2ModelDescribeScheduledInstanceAvailabilityRequest) WithMaxSlotDurationInHours(a int) *ServicesEc2ModelDescribeScheduledInstanceAvailabilityRequest {
	conv_a := javabind.NewGoToJavaInteger()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "withMaxSlotDurationInHours", "com/amazonaws/services/ec2/model/DescribeScheduledInstanceAvailabilityRequest", conv_a.Value().Cast("java/lang/Integer"))
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
	unique_x := &ServicesEc2ModelDescribeScheduledInstanceAvailabilityRequest{}
	unique_x.Callable = dst
	return unique_x
}

// public void setNextToken(java.lang.String)
func (jbobject *ServicesEc2ModelDescribeScheduledInstanceAvailabilityRequest) SetNextToken(a string)  {
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
func (jbobject *ServicesEc2ModelDescribeScheduledInstanceAvailabilityRequest) GetNextToken() string {
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

// public com.amazonaws.services.ec2.model.DescribeScheduledInstanceAvailabilityRequest withNextToken(java.lang.String)
func (jbobject *ServicesEc2ModelDescribeScheduledInstanceAvailabilityRequest) WithNextToken(a string) *ServicesEc2ModelDescribeScheduledInstanceAvailabilityRequest {
	conv_a := javabind.NewGoToJavaString()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "withNextToken", "com/amazonaws/services/ec2/model/DescribeScheduledInstanceAvailabilityRequest", conv_a.Value().Cast("java/lang/String"))
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
	unique_x := &ServicesEc2ModelDescribeScheduledInstanceAvailabilityRequest{}
	unique_x.Callable = dst
	return unique_x
}

// public void setMaxResults(java.lang.Integer)
func (jbobject *ServicesEc2ModelDescribeScheduledInstanceAvailabilityRequest) SetMaxResults(a int)  {
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
func (jbobject *ServicesEc2ModelDescribeScheduledInstanceAvailabilityRequest) GetMaxResults() int {
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

// public com.amazonaws.services.ec2.model.DescribeScheduledInstanceAvailabilityRequest withMaxResults(java.lang.Integer)
func (jbobject *ServicesEc2ModelDescribeScheduledInstanceAvailabilityRequest) WithMaxResults(a int) *ServicesEc2ModelDescribeScheduledInstanceAvailabilityRequest {
	conv_a := javabind.NewGoToJavaInteger()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "withMaxResults", "com/amazonaws/services/ec2/model/DescribeScheduledInstanceAvailabilityRequest", conv_a.Value().Cast("java/lang/Integer"))
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
	unique_x := &ServicesEc2ModelDescribeScheduledInstanceAvailabilityRequest{}
	unique_x.Callable = dst
	return unique_x
}

// public java.util.List<com.amazonaws.services.ec2.model.Filter> getFilters()
func (jbobject *ServicesEc2ModelDescribeScheduledInstanceAvailabilityRequest) GetFilters() []*ServicesEc2ModelFilter {
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
func (jbobject *ServicesEc2ModelDescribeScheduledInstanceAvailabilityRequest) SetFilters(a []*ServicesEc2ModelFilter)  {
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

// public com.amazonaws.services.ec2.model.DescribeScheduledInstanceAvailabilityRequest withFilters(com.amazonaws.services.ec2.model.Filter...)
func (jbobject *ServicesEc2ModelDescribeScheduledInstanceAvailabilityRequest) WithFilters(a ...*ServicesEc2ModelFilter) *ServicesEc2ModelDescribeScheduledInstanceAvailabilityRequest {
	conv_a := javabind.NewGoToJavaObjectArray(javabind.NewGoToJavaCallable(), "com/amazonaws/services/ec2/model/Filter")
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "withFilters", "com/amazonaws/services/ec2/model/DescribeScheduledInstanceAvailabilityRequest", conv_a.Value().Cast("com/amazonaws/services/ec2/model/Filter"))
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
	unique_x := &ServicesEc2ModelDescribeScheduledInstanceAvailabilityRequest{}
	unique_x.Callable = dst
	return unique_x
}

// public com.amazonaws.services.ec2.model.DescribeScheduledInstanceAvailabilityRequest withFilters(java.util.Collection<com.amazonaws.services.ec2.model.Filter>)
func (jbobject *ServicesEc2ModelDescribeScheduledInstanceAvailabilityRequest) WithFilters2(a []*ServicesEc2ModelFilter) *ServicesEc2ModelDescribeScheduledInstanceAvailabilityRequest {
	conv_a := javabind.NewGoToJavaCollection(javabind.NewGoToJavaCallable())
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "withFilters", "com/amazonaws/services/ec2/model/DescribeScheduledInstanceAvailabilityRequest", conv_a.Value().Cast("java/util/Collection"))
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
	unique_x := &ServicesEc2ModelDescribeScheduledInstanceAvailabilityRequest{}
	unique_x.Callable = dst
	return unique_x
}

// public com.amazonaws.Request<com.amazonaws.services.ec2.model.DescribeScheduledInstanceAvailabilityRequest> getDryRunRequest()
func (jbobject *ServicesEc2ModelDescribeScheduledInstanceAvailabilityRequest) GetDryRunRequest() *Request {
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
func (jbobject *ServicesEc2ModelDescribeScheduledInstanceAvailabilityRequest) ToString() string {
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
func (jbobject *ServicesEc2ModelDescribeScheduledInstanceAvailabilityRequest) Equals(a interface{}) bool {
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
func (jbobject *ServicesEc2ModelDescribeScheduledInstanceAvailabilityRequest) HashCode() int {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "hashCode", javabind.Int)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

// public com.amazonaws.services.ec2.model.DescribeScheduledInstanceAvailabilityRequest clone()
func (jbobject *ServicesEc2ModelDescribeScheduledInstanceAvailabilityRequest) Clone3() *ServicesEc2ModelDescribeScheduledInstanceAvailabilityRequest {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "clone", "com/amazonaws/services/ec2/model/DescribeScheduledInstanceAvailabilityRequest")
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
	unique_x := &ServicesEc2ModelDescribeScheduledInstanceAvailabilityRequest{}
	unique_x.Callable = dst
	return unique_x
}

// public com.amazonaws.AmazonWebServiceRequest clone()
func (jbobject *ServicesEc2ModelDescribeScheduledInstanceAvailabilityRequest) Clone() *AmazonWebServiceRequest {
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
func (jbobject *ServicesEc2ModelDescribeScheduledInstanceAvailabilityRequest) Clone2() (*JavaLangObject, error) {
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


