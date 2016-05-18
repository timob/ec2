package ec2

import "github.com/timob/javabind"

type ServicesEc2ModelDescribeSnapshotsRequestInterface interface {
	AmazonWebServiceRequestInterface

	// public java.util.List<java.lang.String> getSnapshotIds()
	GetSnapshotIds() []string

	// public void setSnapshotIds(java.util.Collection<java.lang.String>)
	SetSnapshotIds(a []string) 

	// public com.amazonaws.services.ec2.model.DescribeSnapshotsRequest withSnapshotIds(java.lang.String...)
	WithSnapshotIds(a ...string) *ServicesEc2ModelDescribeSnapshotsRequest

	// public com.amazonaws.services.ec2.model.DescribeSnapshotsRequest withSnapshotIds(java.util.Collection<java.lang.String>)
	WithSnapshotIds2(a []string) *ServicesEc2ModelDescribeSnapshotsRequest

	// public java.util.List<java.lang.String> getOwnerIds()
	GetOwnerIds() []string

	// public void setOwnerIds(java.util.Collection<java.lang.String>)
	SetOwnerIds(a []string) 

	// public com.amazonaws.services.ec2.model.DescribeSnapshotsRequest withOwnerIds(java.lang.String...)
	WithOwnerIds(a ...string) *ServicesEc2ModelDescribeSnapshotsRequest

	// public com.amazonaws.services.ec2.model.DescribeSnapshotsRequest withOwnerIds(java.util.Collection<java.lang.String>)
	WithOwnerIds2(a []string) *ServicesEc2ModelDescribeSnapshotsRequest

	// public java.util.List<java.lang.String> getRestorableByUserIds()
	GetRestorableByUserIds() []string

	// public void setRestorableByUserIds(java.util.Collection<java.lang.String>)
	SetRestorableByUserIds(a []string) 

	// public com.amazonaws.services.ec2.model.DescribeSnapshotsRequest withRestorableByUserIds(java.lang.String...)
	WithRestorableByUserIds(a ...string) *ServicesEc2ModelDescribeSnapshotsRequest

	// public com.amazonaws.services.ec2.model.DescribeSnapshotsRequest withRestorableByUserIds(java.util.Collection<java.lang.String>)
	WithRestorableByUserIds2(a []string) *ServicesEc2ModelDescribeSnapshotsRequest

	// public java.util.List<com.amazonaws.services.ec2.model.Filter> getFilters()
	GetFilters() []*ServicesEc2ModelFilter

	// public void setFilters(java.util.Collection<com.amazonaws.services.ec2.model.Filter>)
	SetFilters(a []*ServicesEc2ModelFilter) 

	// public com.amazonaws.services.ec2.model.DescribeSnapshotsRequest withFilters(com.amazonaws.services.ec2.model.Filter...)
	WithFilters(a ...*ServicesEc2ModelFilter) *ServicesEc2ModelDescribeSnapshotsRequest

	// public com.amazonaws.services.ec2.model.DescribeSnapshotsRequest withFilters(java.util.Collection<com.amazonaws.services.ec2.model.Filter>)
	WithFilters2(a []*ServicesEc2ModelFilter) *ServicesEc2ModelDescribeSnapshotsRequest

	// public void setNextToken(java.lang.String)
	SetNextToken(a string) 

	// public java.lang.String getNextToken()
	GetNextToken() string

	// public com.amazonaws.services.ec2.model.DescribeSnapshotsRequest withNextToken(java.lang.String)
	WithNextToken(a string) *ServicesEc2ModelDescribeSnapshotsRequest

	// public void setMaxResults(java.lang.Integer)
	SetMaxResults(a int) 

	// public java.lang.Integer getMaxResults()
	GetMaxResults() int

	// public com.amazonaws.services.ec2.model.DescribeSnapshotsRequest withMaxResults(java.lang.Integer)
	WithMaxResults(a int) *ServicesEc2ModelDescribeSnapshotsRequest

	// public com.amazonaws.Request<com.amazonaws.services.ec2.model.DescribeSnapshotsRequest> getDryRunRequest()
	GetDryRunRequest() *Request

	// public com.amazonaws.services.ec2.model.DescribeSnapshotsRequest clone()
	Clone3() *ServicesEc2ModelDescribeSnapshotsRequest
}

type ServicesEc2ModelDescribeSnapshotsRequest struct {
	AmazonWebServiceRequest
}

// public com.amazonaws.services.ec2.model.DescribeSnapshotsRequest()
func NewServicesEc2ModelDescribeSnapshotsRequest() (*ServicesEc2ModelDescribeSnapshotsRequest) {

	obj, err := javabind.GetEnv().NewObject("com/amazonaws/services/ec2/model/DescribeSnapshotsRequest")
	if err != nil {
		panic(err)
	}
	x := &ServicesEc2ModelDescribeSnapshotsRequest{}
	x.Callable = &javabind.Callable{obj}
	return x
}

// public java.util.List<java.lang.String> getSnapshotIds()
func (jbobject *ServicesEc2ModelDescribeSnapshotsRequest) GetSnapshotIds() []string {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getSnapshotIds", "java/util/List")
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

// public void setSnapshotIds(java.util.Collection<java.lang.String>)
func (jbobject *ServicesEc2ModelDescribeSnapshotsRequest) SetSnapshotIds(a []string)  {
	conv_a := javabind.NewGoToJavaCollection(javabind.NewGoToJavaString())
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "setSnapshotIds", javabind.Void, conv_a.Value().Cast("java/util/Collection"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()

}

// public com.amazonaws.services.ec2.model.DescribeSnapshotsRequest withSnapshotIds(java.lang.String...)
func (jbobject *ServicesEc2ModelDescribeSnapshotsRequest) WithSnapshotIds(a ...string) *ServicesEc2ModelDescribeSnapshotsRequest {
	conv_a := javabind.NewGoToJavaObjectArray(javabind.NewGoToJavaString(), "java/lang/String")
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "withSnapshotIds", "com/amazonaws/services/ec2/model/DescribeSnapshotsRequest", conv_a.Value().Cast("java/lang/String"))
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
	unique_x := &ServicesEc2ModelDescribeSnapshotsRequest{}
	unique_x.Callable = dst
	return unique_x
}

// public com.amazonaws.services.ec2.model.DescribeSnapshotsRequest withSnapshotIds(java.util.Collection<java.lang.String>)
func (jbobject *ServicesEc2ModelDescribeSnapshotsRequest) WithSnapshotIds2(a []string) *ServicesEc2ModelDescribeSnapshotsRequest {
	conv_a := javabind.NewGoToJavaCollection(javabind.NewGoToJavaString())
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "withSnapshotIds", "com/amazonaws/services/ec2/model/DescribeSnapshotsRequest", conv_a.Value().Cast("java/util/Collection"))
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
	unique_x := &ServicesEc2ModelDescribeSnapshotsRequest{}
	unique_x.Callable = dst
	return unique_x
}

// public java.util.List<java.lang.String> getOwnerIds()
func (jbobject *ServicesEc2ModelDescribeSnapshotsRequest) GetOwnerIds() []string {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getOwnerIds", "java/util/List")
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

// public void setOwnerIds(java.util.Collection<java.lang.String>)
func (jbobject *ServicesEc2ModelDescribeSnapshotsRequest) SetOwnerIds(a []string)  {
	conv_a := javabind.NewGoToJavaCollection(javabind.NewGoToJavaString())
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "setOwnerIds", javabind.Void, conv_a.Value().Cast("java/util/Collection"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()

}

// public com.amazonaws.services.ec2.model.DescribeSnapshotsRequest withOwnerIds(java.lang.String...)
func (jbobject *ServicesEc2ModelDescribeSnapshotsRequest) WithOwnerIds(a ...string) *ServicesEc2ModelDescribeSnapshotsRequest {
	conv_a := javabind.NewGoToJavaObjectArray(javabind.NewGoToJavaString(), "java/lang/String")
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "withOwnerIds", "com/amazonaws/services/ec2/model/DescribeSnapshotsRequest", conv_a.Value().Cast("java/lang/String"))
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
	unique_x := &ServicesEc2ModelDescribeSnapshotsRequest{}
	unique_x.Callable = dst
	return unique_x
}

// public com.amazonaws.services.ec2.model.DescribeSnapshotsRequest withOwnerIds(java.util.Collection<java.lang.String>)
func (jbobject *ServicesEc2ModelDescribeSnapshotsRequest) WithOwnerIds2(a []string) *ServicesEc2ModelDescribeSnapshotsRequest {
	conv_a := javabind.NewGoToJavaCollection(javabind.NewGoToJavaString())
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "withOwnerIds", "com/amazonaws/services/ec2/model/DescribeSnapshotsRequest", conv_a.Value().Cast("java/util/Collection"))
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
	unique_x := &ServicesEc2ModelDescribeSnapshotsRequest{}
	unique_x.Callable = dst
	return unique_x
}

// public java.util.List<java.lang.String> getRestorableByUserIds()
func (jbobject *ServicesEc2ModelDescribeSnapshotsRequest) GetRestorableByUserIds() []string {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getRestorableByUserIds", "java/util/List")
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

// public void setRestorableByUserIds(java.util.Collection<java.lang.String>)
func (jbobject *ServicesEc2ModelDescribeSnapshotsRequest) SetRestorableByUserIds(a []string)  {
	conv_a := javabind.NewGoToJavaCollection(javabind.NewGoToJavaString())
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "setRestorableByUserIds", javabind.Void, conv_a.Value().Cast("java/util/Collection"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()

}

// public com.amazonaws.services.ec2.model.DescribeSnapshotsRequest withRestorableByUserIds(java.lang.String...)
func (jbobject *ServicesEc2ModelDescribeSnapshotsRequest) WithRestorableByUserIds(a ...string) *ServicesEc2ModelDescribeSnapshotsRequest {
	conv_a := javabind.NewGoToJavaObjectArray(javabind.NewGoToJavaString(), "java/lang/String")
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "withRestorableByUserIds", "com/amazonaws/services/ec2/model/DescribeSnapshotsRequest", conv_a.Value().Cast("java/lang/String"))
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
	unique_x := &ServicesEc2ModelDescribeSnapshotsRequest{}
	unique_x.Callable = dst
	return unique_x
}

// public com.amazonaws.services.ec2.model.DescribeSnapshotsRequest withRestorableByUserIds(java.util.Collection<java.lang.String>)
func (jbobject *ServicesEc2ModelDescribeSnapshotsRequest) WithRestorableByUserIds2(a []string) *ServicesEc2ModelDescribeSnapshotsRequest {
	conv_a := javabind.NewGoToJavaCollection(javabind.NewGoToJavaString())
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "withRestorableByUserIds", "com/amazonaws/services/ec2/model/DescribeSnapshotsRequest", conv_a.Value().Cast("java/util/Collection"))
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
	unique_x := &ServicesEc2ModelDescribeSnapshotsRequest{}
	unique_x.Callable = dst
	return unique_x
}

// public java.util.List<com.amazonaws.services.ec2.model.Filter> getFilters()
func (jbobject *ServicesEc2ModelDescribeSnapshotsRequest) GetFilters() []*ServicesEc2ModelFilter {
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
func (jbobject *ServicesEc2ModelDescribeSnapshotsRequest) SetFilters(a []*ServicesEc2ModelFilter)  {
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

// public com.amazonaws.services.ec2.model.DescribeSnapshotsRequest withFilters(com.amazonaws.services.ec2.model.Filter...)
func (jbobject *ServicesEc2ModelDescribeSnapshotsRequest) WithFilters(a ...*ServicesEc2ModelFilter) *ServicesEc2ModelDescribeSnapshotsRequest {
	conv_a := javabind.NewGoToJavaObjectArray(javabind.NewGoToJavaCallable(), "com/amazonaws/services/ec2/model/Filter")
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "withFilters", "com/amazonaws/services/ec2/model/DescribeSnapshotsRequest", conv_a.Value().Cast("com/amazonaws/services/ec2/model/Filter"))
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
	unique_x := &ServicesEc2ModelDescribeSnapshotsRequest{}
	unique_x.Callable = dst
	return unique_x
}

// public com.amazonaws.services.ec2.model.DescribeSnapshotsRequest withFilters(java.util.Collection<com.amazonaws.services.ec2.model.Filter>)
func (jbobject *ServicesEc2ModelDescribeSnapshotsRequest) WithFilters2(a []*ServicesEc2ModelFilter) *ServicesEc2ModelDescribeSnapshotsRequest {
	conv_a := javabind.NewGoToJavaCollection(javabind.NewGoToJavaCallable())
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "withFilters", "com/amazonaws/services/ec2/model/DescribeSnapshotsRequest", conv_a.Value().Cast("java/util/Collection"))
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
	unique_x := &ServicesEc2ModelDescribeSnapshotsRequest{}
	unique_x.Callable = dst
	return unique_x
}

// public void setNextToken(java.lang.String)
func (jbobject *ServicesEc2ModelDescribeSnapshotsRequest) SetNextToken(a string)  {
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
func (jbobject *ServicesEc2ModelDescribeSnapshotsRequest) GetNextToken() string {
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

// public com.amazonaws.services.ec2.model.DescribeSnapshotsRequest withNextToken(java.lang.String)
func (jbobject *ServicesEc2ModelDescribeSnapshotsRequest) WithNextToken(a string) *ServicesEc2ModelDescribeSnapshotsRequest {
	conv_a := javabind.NewGoToJavaString()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "withNextToken", "com/amazonaws/services/ec2/model/DescribeSnapshotsRequest", conv_a.Value().Cast("java/lang/String"))
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
	unique_x := &ServicesEc2ModelDescribeSnapshotsRequest{}
	unique_x.Callable = dst
	return unique_x
}

// public void setMaxResults(java.lang.Integer)
func (jbobject *ServicesEc2ModelDescribeSnapshotsRequest) SetMaxResults(a int)  {
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
func (jbobject *ServicesEc2ModelDescribeSnapshotsRequest) GetMaxResults() int {
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

// public com.amazonaws.services.ec2.model.DescribeSnapshotsRequest withMaxResults(java.lang.Integer)
func (jbobject *ServicesEc2ModelDescribeSnapshotsRequest) WithMaxResults(a int) *ServicesEc2ModelDescribeSnapshotsRequest {
	conv_a := javabind.NewGoToJavaInteger()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "withMaxResults", "com/amazonaws/services/ec2/model/DescribeSnapshotsRequest", conv_a.Value().Cast("java/lang/Integer"))
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
	unique_x := &ServicesEc2ModelDescribeSnapshotsRequest{}
	unique_x.Callable = dst
	return unique_x
}

// public com.amazonaws.Request<com.amazonaws.services.ec2.model.DescribeSnapshotsRequest> getDryRunRequest()
func (jbobject *ServicesEc2ModelDescribeSnapshotsRequest) GetDryRunRequest() *Request {
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
func (jbobject *ServicesEc2ModelDescribeSnapshotsRequest) ToString() string {
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
func (jbobject *ServicesEc2ModelDescribeSnapshotsRequest) Equals(a interface{}) bool {
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
func (jbobject *ServicesEc2ModelDescribeSnapshotsRequest) HashCode() int {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "hashCode", javabind.Int)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

// public com.amazonaws.services.ec2.model.DescribeSnapshotsRequest clone()
func (jbobject *ServicesEc2ModelDescribeSnapshotsRequest) Clone3() *ServicesEc2ModelDescribeSnapshotsRequest {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "clone", "com/amazonaws/services/ec2/model/DescribeSnapshotsRequest")
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
	unique_x := &ServicesEc2ModelDescribeSnapshotsRequest{}
	unique_x.Callable = dst
	return unique_x
}

// public com.amazonaws.AmazonWebServiceRequest clone()
func (jbobject *ServicesEc2ModelDescribeSnapshotsRequest) Clone() *AmazonWebServiceRequest {
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
func (jbobject *ServicesEc2ModelDescribeSnapshotsRequest) Clone2() (*JavaLangObject, error) {
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


