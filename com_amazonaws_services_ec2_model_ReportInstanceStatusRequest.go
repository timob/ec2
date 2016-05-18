package ec2

import "time"
import "github.com/timob/javabind"

type ServicesEc2ModelReportInstanceStatusRequestInterface interface {
	AmazonWebServiceRequestInterface

	// public java.util.List<java.lang.String> getInstances()
	GetInstances() []string

	// public void setInstances(java.util.Collection<java.lang.String>)
	SetInstances(a []string) 

	// public com.amazonaws.services.ec2.model.ReportInstanceStatusRequest withInstances(java.lang.String...)
	WithInstances(a ...string) *ServicesEc2ModelReportInstanceStatusRequest

	// public com.amazonaws.services.ec2.model.ReportInstanceStatusRequest withInstances(java.util.Collection<java.lang.String>)
	WithInstances2(a []string) *ServicesEc2ModelReportInstanceStatusRequest

	// public void setStatus(java.lang.String)
	SetStatus2(a string) 

	// public java.lang.String getStatus()
	GetStatus() string

	// public com.amazonaws.services.ec2.model.ReportInstanceStatusRequest withStatus(java.lang.String)
	WithStatus2(a string) *ServicesEc2ModelReportInstanceStatusRequest

	// public void setStatus(com.amazonaws.services.ec2.model.ReportStatusType)
	SetStatus(a ServicesEc2ModelReportStatusTypeInterface) 

	// public com.amazonaws.services.ec2.model.ReportInstanceStatusRequest withStatus(com.amazonaws.services.ec2.model.ReportStatusType)
	WithStatus(a ServicesEc2ModelReportStatusTypeInterface) *ServicesEc2ModelReportInstanceStatusRequest

	// public void setStartTime(java.util.Date)
	SetStartTime(a time.Time) 

	// public java.util.Date getStartTime()
	GetStartTime() time.Time

	// public com.amazonaws.services.ec2.model.ReportInstanceStatusRequest withStartTime(java.util.Date)
	WithStartTime(a time.Time) *ServicesEc2ModelReportInstanceStatusRequest

	// public void setEndTime(java.util.Date)
	SetEndTime(a time.Time) 

	// public java.util.Date getEndTime()
	GetEndTime() time.Time

	// public com.amazonaws.services.ec2.model.ReportInstanceStatusRequest withEndTime(java.util.Date)
	WithEndTime(a time.Time) *ServicesEc2ModelReportInstanceStatusRequest

	// public java.util.List<java.lang.String> getReasonCodes()
	GetReasonCodes() []string

	// public void setReasonCodes(java.util.Collection<java.lang.String>)
	SetReasonCodes(a []string) 

	// public com.amazonaws.services.ec2.model.ReportInstanceStatusRequest withReasonCodes(java.lang.String...)
	WithReasonCodes2(a ...string) *ServicesEc2ModelReportInstanceStatusRequest

	// public com.amazonaws.services.ec2.model.ReportInstanceStatusRequest withReasonCodes(java.util.Collection<java.lang.String>)
	WithReasonCodes3(a []string) *ServicesEc2ModelReportInstanceStatusRequest

	// public com.amazonaws.services.ec2.model.ReportInstanceStatusRequest withReasonCodes(com.amazonaws.services.ec2.model.ReportInstanceReasonCodes...)
	WithReasonCodes(a ...*ServicesEc2ModelReportInstanceReasonCodes) *ServicesEc2ModelReportInstanceStatusRequest

	// public void setDescription(java.lang.String)
	SetDescription(a string) 

	// public java.lang.String getDescription()
	GetDescription() string

	// public com.amazonaws.services.ec2.model.ReportInstanceStatusRequest withDescription(java.lang.String)
	WithDescription(a string) *ServicesEc2ModelReportInstanceStatusRequest

	// public com.amazonaws.Request<com.amazonaws.services.ec2.model.ReportInstanceStatusRequest> getDryRunRequest()
	GetDryRunRequest() *Request

	// public com.amazonaws.services.ec2.model.ReportInstanceStatusRequest clone()
	Clone3() *ServicesEc2ModelReportInstanceStatusRequest
}

type ServicesEc2ModelReportInstanceStatusRequest struct {
	AmazonWebServiceRequest
}

// public com.amazonaws.services.ec2.model.ReportInstanceStatusRequest()
func NewServicesEc2ModelReportInstanceStatusRequest() (*ServicesEc2ModelReportInstanceStatusRequest) {

	obj, err := javabind.GetEnv().NewObject("com/amazonaws/services/ec2/model/ReportInstanceStatusRequest")
	if err != nil {
		panic(err)
	}
	x := &ServicesEc2ModelReportInstanceStatusRequest{}
	x.Callable = &javabind.Callable{obj}
	return x
}

// public java.util.List<java.lang.String> getInstances()
func (jbobject *ServicesEc2ModelReportInstanceStatusRequest) GetInstances() []string {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getInstances", "java/util/List")
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

// public void setInstances(java.util.Collection<java.lang.String>)
func (jbobject *ServicesEc2ModelReportInstanceStatusRequest) SetInstances(a []string)  {
	conv_a := javabind.NewGoToJavaCollection(javabind.NewGoToJavaString())
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "setInstances", javabind.Void, conv_a.Value().Cast("java/util/Collection"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()

}

// public com.amazonaws.services.ec2.model.ReportInstanceStatusRequest withInstances(java.lang.String...)
func (jbobject *ServicesEc2ModelReportInstanceStatusRequest) WithInstances(a ...string) *ServicesEc2ModelReportInstanceStatusRequest {
	conv_a := javabind.NewGoToJavaObjectArray(javabind.NewGoToJavaString(), "java/lang/String")
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "withInstances", "com/amazonaws/services/ec2/model/ReportInstanceStatusRequest", conv_a.Value().Cast("java/lang/String"))
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
	unique_x := &ServicesEc2ModelReportInstanceStatusRequest{}
	unique_x.Callable = dst
	return unique_x
}

// public com.amazonaws.services.ec2.model.ReportInstanceStatusRequest withInstances(java.util.Collection<java.lang.String>)
func (jbobject *ServicesEc2ModelReportInstanceStatusRequest) WithInstances2(a []string) *ServicesEc2ModelReportInstanceStatusRequest {
	conv_a := javabind.NewGoToJavaCollection(javabind.NewGoToJavaString())
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "withInstances", "com/amazonaws/services/ec2/model/ReportInstanceStatusRequest", conv_a.Value().Cast("java/util/Collection"))
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
	unique_x := &ServicesEc2ModelReportInstanceStatusRequest{}
	unique_x.Callable = dst
	return unique_x
}

// public void setStatus(java.lang.String)
func (jbobject *ServicesEc2ModelReportInstanceStatusRequest) SetStatus2(a string)  {
	conv_a := javabind.NewGoToJavaString()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "setStatus", javabind.Void, conv_a.Value().Cast("java/lang/String"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()

}

// public java.lang.String getStatus()
func (jbobject *ServicesEc2ModelReportInstanceStatusRequest) GetStatus() string {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getStatus", "java/lang/String")
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

// public com.amazonaws.services.ec2.model.ReportInstanceStatusRequest withStatus(java.lang.String)
func (jbobject *ServicesEc2ModelReportInstanceStatusRequest) WithStatus2(a string) *ServicesEc2ModelReportInstanceStatusRequest {
	conv_a := javabind.NewGoToJavaString()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "withStatus", "com/amazonaws/services/ec2/model/ReportInstanceStatusRequest", conv_a.Value().Cast("java/lang/String"))
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
	unique_x := &ServicesEc2ModelReportInstanceStatusRequest{}
	unique_x.Callable = dst
	return unique_x
}

// public void setStatus(com.amazonaws.services.ec2.model.ReportStatusType)
func (jbobject *ServicesEc2ModelReportInstanceStatusRequest) SetStatus(a ServicesEc2ModelReportStatusTypeInterface)  {
	conv_a := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "setStatus", javabind.Void, conv_a.Value().Cast("com/amazonaws/services/ec2/model/ReportStatusType"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()

}

// public com.amazonaws.services.ec2.model.ReportInstanceStatusRequest withStatus(com.amazonaws.services.ec2.model.ReportStatusType)
func (jbobject *ServicesEc2ModelReportInstanceStatusRequest) WithStatus(a ServicesEc2ModelReportStatusTypeInterface) *ServicesEc2ModelReportInstanceStatusRequest {
	conv_a := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "withStatus", "com/amazonaws/services/ec2/model/ReportInstanceStatusRequest", conv_a.Value().Cast("com/amazonaws/services/ec2/model/ReportStatusType"))
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
	unique_x := &ServicesEc2ModelReportInstanceStatusRequest{}
	unique_x.Callable = dst
	return unique_x
}

// public void setStartTime(java.util.Date)
func (jbobject *ServicesEc2ModelReportInstanceStatusRequest) SetStartTime(a time.Time)  {
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
func (jbobject *ServicesEc2ModelReportInstanceStatusRequest) GetStartTime() time.Time {
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

// public com.amazonaws.services.ec2.model.ReportInstanceStatusRequest withStartTime(java.util.Date)
func (jbobject *ServicesEc2ModelReportInstanceStatusRequest) WithStartTime(a time.Time) *ServicesEc2ModelReportInstanceStatusRequest {
	conv_a := javabind.NewGoToJavaDate()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "withStartTime", "com/amazonaws/services/ec2/model/ReportInstanceStatusRequest", conv_a.Value().Cast("java/util/Date"))
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
	unique_x := &ServicesEc2ModelReportInstanceStatusRequest{}
	unique_x.Callable = dst
	return unique_x
}

// public void setEndTime(java.util.Date)
func (jbobject *ServicesEc2ModelReportInstanceStatusRequest) SetEndTime(a time.Time)  {
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
func (jbobject *ServicesEc2ModelReportInstanceStatusRequest) GetEndTime() time.Time {
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

// public com.amazonaws.services.ec2.model.ReportInstanceStatusRequest withEndTime(java.util.Date)
func (jbobject *ServicesEc2ModelReportInstanceStatusRequest) WithEndTime(a time.Time) *ServicesEc2ModelReportInstanceStatusRequest {
	conv_a := javabind.NewGoToJavaDate()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "withEndTime", "com/amazonaws/services/ec2/model/ReportInstanceStatusRequest", conv_a.Value().Cast("java/util/Date"))
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
	unique_x := &ServicesEc2ModelReportInstanceStatusRequest{}
	unique_x.Callable = dst
	return unique_x
}

// public java.util.List<java.lang.String> getReasonCodes()
func (jbobject *ServicesEc2ModelReportInstanceStatusRequest) GetReasonCodes() []string {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getReasonCodes", "java/util/List")
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

// public void setReasonCodes(java.util.Collection<java.lang.String>)
func (jbobject *ServicesEc2ModelReportInstanceStatusRequest) SetReasonCodes(a []string)  {
	conv_a := javabind.NewGoToJavaCollection(javabind.NewGoToJavaString())
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "setReasonCodes", javabind.Void, conv_a.Value().Cast("java/util/Collection"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()

}

// public com.amazonaws.services.ec2.model.ReportInstanceStatusRequest withReasonCodes(java.lang.String...)
func (jbobject *ServicesEc2ModelReportInstanceStatusRequest) WithReasonCodes2(a ...string) *ServicesEc2ModelReportInstanceStatusRequest {
	conv_a := javabind.NewGoToJavaObjectArray(javabind.NewGoToJavaString(), "java/lang/String")
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "withReasonCodes", "com/amazonaws/services/ec2/model/ReportInstanceStatusRequest", conv_a.Value().Cast("java/lang/String"))
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
	unique_x := &ServicesEc2ModelReportInstanceStatusRequest{}
	unique_x.Callable = dst
	return unique_x
}

// public com.amazonaws.services.ec2.model.ReportInstanceStatusRequest withReasonCodes(java.util.Collection<java.lang.String>)
func (jbobject *ServicesEc2ModelReportInstanceStatusRequest) WithReasonCodes3(a []string) *ServicesEc2ModelReportInstanceStatusRequest {
	conv_a := javabind.NewGoToJavaCollection(javabind.NewGoToJavaString())
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "withReasonCodes", "com/amazonaws/services/ec2/model/ReportInstanceStatusRequest", conv_a.Value().Cast("java/util/Collection"))
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
	unique_x := &ServicesEc2ModelReportInstanceStatusRequest{}
	unique_x.Callable = dst
	return unique_x
}

// public com.amazonaws.services.ec2.model.ReportInstanceStatusRequest withReasonCodes(com.amazonaws.services.ec2.model.ReportInstanceReasonCodes...)
func (jbobject *ServicesEc2ModelReportInstanceStatusRequest) WithReasonCodes(a ...*ServicesEc2ModelReportInstanceReasonCodes) *ServicesEc2ModelReportInstanceStatusRequest {
	conv_a := javabind.NewGoToJavaObjectArray(javabind.NewGoToJavaCallable(), "com/amazonaws/services/ec2/model/ReportInstanceReasonCodes")
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "withReasonCodes", "com/amazonaws/services/ec2/model/ReportInstanceStatusRequest", conv_a.Value().Cast("com/amazonaws/services/ec2/model/ReportInstanceReasonCodes"))
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
	unique_x := &ServicesEc2ModelReportInstanceStatusRequest{}
	unique_x.Callable = dst
	return unique_x
}

// public void setDescription(java.lang.String)
func (jbobject *ServicesEc2ModelReportInstanceStatusRequest) SetDescription(a string)  {
	conv_a := javabind.NewGoToJavaString()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "setDescription", javabind.Void, conv_a.Value().Cast("java/lang/String"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()

}

// public java.lang.String getDescription()
func (jbobject *ServicesEc2ModelReportInstanceStatusRequest) GetDescription() string {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getDescription", "java/lang/String")
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

// public com.amazonaws.services.ec2.model.ReportInstanceStatusRequest withDescription(java.lang.String)
func (jbobject *ServicesEc2ModelReportInstanceStatusRequest) WithDescription(a string) *ServicesEc2ModelReportInstanceStatusRequest {
	conv_a := javabind.NewGoToJavaString()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "withDescription", "com/amazonaws/services/ec2/model/ReportInstanceStatusRequest", conv_a.Value().Cast("java/lang/String"))
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
	unique_x := &ServicesEc2ModelReportInstanceStatusRequest{}
	unique_x.Callable = dst
	return unique_x
}

// public com.amazonaws.Request<com.amazonaws.services.ec2.model.ReportInstanceStatusRequest> getDryRunRequest()
func (jbobject *ServicesEc2ModelReportInstanceStatusRequest) GetDryRunRequest() *Request {
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
func (jbobject *ServicesEc2ModelReportInstanceStatusRequest) ToString() string {
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
func (jbobject *ServicesEc2ModelReportInstanceStatusRequest) Equals(a interface{}) bool {
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
func (jbobject *ServicesEc2ModelReportInstanceStatusRequest) HashCode() int {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "hashCode", javabind.Int)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

// public com.amazonaws.services.ec2.model.ReportInstanceStatusRequest clone()
func (jbobject *ServicesEc2ModelReportInstanceStatusRequest) Clone3() *ServicesEc2ModelReportInstanceStatusRequest {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "clone", "com/amazonaws/services/ec2/model/ReportInstanceStatusRequest")
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
	unique_x := &ServicesEc2ModelReportInstanceStatusRequest{}
	unique_x.Callable = dst
	return unique_x
}

// public com.amazonaws.AmazonWebServiceRequest clone()
func (jbobject *ServicesEc2ModelReportInstanceStatusRequest) Clone() *AmazonWebServiceRequest {
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
func (jbobject *ServicesEc2ModelReportInstanceStatusRequest) Clone2() (*JavaLangObject, error) {
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


