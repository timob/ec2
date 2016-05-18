package ec2

import "time"
import "github.com/timob/javabind"

type ServicesEc2ModelReservedInstancesModificationInterface interface {
	JavaLangObjectInterface

	// public void setReservedInstancesModificationId(java.lang.String)
	SetReservedInstancesModificationId(a string) 

	// public java.lang.String getReservedInstancesModificationId()
	GetReservedInstancesModificationId() string

	// public com.amazonaws.services.ec2.model.ReservedInstancesModification withReservedInstancesModificationId(java.lang.String)
	WithReservedInstancesModificationId(a string) *ServicesEc2ModelReservedInstancesModification

	// public java.util.List<com.amazonaws.services.ec2.model.ReservedInstancesId> getReservedInstancesIds()
	GetReservedInstancesIds() []*ServicesEc2ModelReservedInstancesId

	// public void setReservedInstancesIds(java.util.Collection<com.amazonaws.services.ec2.model.ReservedInstancesId>)
	SetReservedInstancesIds(a []*ServicesEc2ModelReservedInstancesId) 

	// public com.amazonaws.services.ec2.model.ReservedInstancesModification withReservedInstancesIds(com.amazonaws.services.ec2.model.ReservedInstancesId...)
	WithReservedInstancesIds(a ...*ServicesEc2ModelReservedInstancesId) *ServicesEc2ModelReservedInstancesModification

	// public com.amazonaws.services.ec2.model.ReservedInstancesModification withReservedInstancesIds(java.util.Collection<com.amazonaws.services.ec2.model.ReservedInstancesId>)
	WithReservedInstancesIds2(a []*ServicesEc2ModelReservedInstancesId) *ServicesEc2ModelReservedInstancesModification

	// public java.util.List<com.amazonaws.services.ec2.model.ReservedInstancesModificationResult> getModificationResults()
	GetModificationResults() []*ServicesEc2ModelReservedInstancesModificationResult

	// public void setModificationResults(java.util.Collection<com.amazonaws.services.ec2.model.ReservedInstancesModificationResult>)
	SetModificationResults(a []*ServicesEc2ModelReservedInstancesModificationResult) 

	// public com.amazonaws.services.ec2.model.ReservedInstancesModification withModificationResults(com.amazonaws.services.ec2.model.ReservedInstancesModificationResult...)
	WithModificationResults(a ...*ServicesEc2ModelReservedInstancesModificationResult) *ServicesEc2ModelReservedInstancesModification

	// public com.amazonaws.services.ec2.model.ReservedInstancesModification withModificationResults(java.util.Collection<com.amazonaws.services.ec2.model.ReservedInstancesModificationResult>)
	WithModificationResults2(a []*ServicesEc2ModelReservedInstancesModificationResult) *ServicesEc2ModelReservedInstancesModification

	// public void setCreateDate(java.util.Date)
	SetCreateDate(a time.Time) 

	// public java.util.Date getCreateDate()
	GetCreateDate() time.Time

	// public com.amazonaws.services.ec2.model.ReservedInstancesModification withCreateDate(java.util.Date)
	WithCreateDate(a time.Time) *ServicesEc2ModelReservedInstancesModification

	// public void setUpdateDate(java.util.Date)
	SetUpdateDate(a time.Time) 

	// public java.util.Date getUpdateDate()
	GetUpdateDate() time.Time

	// public com.amazonaws.services.ec2.model.ReservedInstancesModification withUpdateDate(java.util.Date)
	WithUpdateDate(a time.Time) *ServicesEc2ModelReservedInstancesModification

	// public void setEffectiveDate(java.util.Date)
	SetEffectiveDate(a time.Time) 

	// public java.util.Date getEffectiveDate()
	GetEffectiveDate() time.Time

	// public com.amazonaws.services.ec2.model.ReservedInstancesModification withEffectiveDate(java.util.Date)
	WithEffectiveDate(a time.Time) *ServicesEc2ModelReservedInstancesModification

	// public void setStatus(java.lang.String)
	SetStatus(a string) 

	// public java.lang.String getStatus()
	GetStatus() string

	// public com.amazonaws.services.ec2.model.ReservedInstancesModification withStatus(java.lang.String)
	WithStatus(a string) *ServicesEc2ModelReservedInstancesModification

	// public void setStatusMessage(java.lang.String)
	SetStatusMessage(a string) 

	// public java.lang.String getStatusMessage()
	GetStatusMessage() string

	// public com.amazonaws.services.ec2.model.ReservedInstancesModification withStatusMessage(java.lang.String)
	WithStatusMessage(a string) *ServicesEc2ModelReservedInstancesModification

	// public void setClientToken(java.lang.String)
	SetClientToken(a string) 

	// public java.lang.String getClientToken()
	GetClientToken() string

	// public com.amazonaws.services.ec2.model.ReservedInstancesModification withClientToken(java.lang.String)
	WithClientToken(a string) *ServicesEc2ModelReservedInstancesModification

	// public com.amazonaws.services.ec2.model.ReservedInstancesModification clone()
	Clone() *ServicesEc2ModelReservedInstancesModification
}

type ServicesEc2ModelReservedInstancesModification struct {
	JavaLangObject
}

// public com.amazonaws.services.ec2.model.ReservedInstancesModification()
func NewServicesEc2ModelReservedInstancesModification() (*ServicesEc2ModelReservedInstancesModification) {

	obj, err := javabind.GetEnv().NewObject("com/amazonaws/services/ec2/model/ReservedInstancesModification")
	if err != nil {
		panic(err)
	}
	x := &ServicesEc2ModelReservedInstancesModification{}
	x.Callable = &javabind.Callable{obj}
	return x
}

// public void setReservedInstancesModificationId(java.lang.String)
func (jbobject *ServicesEc2ModelReservedInstancesModification) SetReservedInstancesModificationId(a string)  {
	conv_a := javabind.NewGoToJavaString()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "setReservedInstancesModificationId", javabind.Void, conv_a.Value().Cast("java/lang/String"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()

}

// public java.lang.String getReservedInstancesModificationId()
func (jbobject *ServicesEc2ModelReservedInstancesModification) GetReservedInstancesModificationId() string {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getReservedInstancesModificationId", "java/lang/String")
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

// public com.amazonaws.services.ec2.model.ReservedInstancesModification withReservedInstancesModificationId(java.lang.String)
func (jbobject *ServicesEc2ModelReservedInstancesModification) WithReservedInstancesModificationId(a string) *ServicesEc2ModelReservedInstancesModification {
	conv_a := javabind.NewGoToJavaString()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "withReservedInstancesModificationId", "com/amazonaws/services/ec2/model/ReservedInstancesModification", conv_a.Value().Cast("java/lang/String"))
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
	unique_x := &ServicesEc2ModelReservedInstancesModification{}
	unique_x.Callable = dst
	return unique_x
}

// public java.util.List<com.amazonaws.services.ec2.model.ReservedInstancesId> getReservedInstancesIds()
func (jbobject *ServicesEc2ModelReservedInstancesModification) GetReservedInstancesIds() []*ServicesEc2ModelReservedInstancesId {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getReservedInstancesIds", "java/util/List")
	if err != nil {
		panic(err)
	}
	retconv := javabind.NewJavaToGoList(javabind.NewJavaToGoCallable())
	dst := new([]*ServicesEc2ModelReservedInstancesId)
	retconv.Dest(dst)
	if err := retconv.Convert(javabind.ObjectRef(jret)); err != nil {
		panic(err)
	}
	retconv.CleanUp()
	return *dst
}

// public void setReservedInstancesIds(java.util.Collection<com.amazonaws.services.ec2.model.ReservedInstancesId>)
func (jbobject *ServicesEc2ModelReservedInstancesModification) SetReservedInstancesIds(a []*ServicesEc2ModelReservedInstancesId)  {
	conv_a := javabind.NewGoToJavaCollection(javabind.NewGoToJavaCallable())
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "setReservedInstancesIds", javabind.Void, conv_a.Value().Cast("java/util/Collection"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()

}

// public com.amazonaws.services.ec2.model.ReservedInstancesModification withReservedInstancesIds(com.amazonaws.services.ec2.model.ReservedInstancesId...)
func (jbobject *ServicesEc2ModelReservedInstancesModification) WithReservedInstancesIds(a ...*ServicesEc2ModelReservedInstancesId) *ServicesEc2ModelReservedInstancesModification {
	conv_a := javabind.NewGoToJavaObjectArray(javabind.NewGoToJavaCallable(), "com/amazonaws/services/ec2/model/ReservedInstancesId")
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "withReservedInstancesIds", "com/amazonaws/services/ec2/model/ReservedInstancesModification", conv_a.Value().Cast("com/amazonaws/services/ec2/model/ReservedInstancesId"))
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
	unique_x := &ServicesEc2ModelReservedInstancesModification{}
	unique_x.Callable = dst
	return unique_x
}

// public com.amazonaws.services.ec2.model.ReservedInstancesModification withReservedInstancesIds(java.util.Collection<com.amazonaws.services.ec2.model.ReservedInstancesId>)
func (jbobject *ServicesEc2ModelReservedInstancesModification) WithReservedInstancesIds2(a []*ServicesEc2ModelReservedInstancesId) *ServicesEc2ModelReservedInstancesModification {
	conv_a := javabind.NewGoToJavaCollection(javabind.NewGoToJavaCallable())
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "withReservedInstancesIds", "com/amazonaws/services/ec2/model/ReservedInstancesModification", conv_a.Value().Cast("java/util/Collection"))
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
	unique_x := &ServicesEc2ModelReservedInstancesModification{}
	unique_x.Callable = dst
	return unique_x
}

// public java.util.List<com.amazonaws.services.ec2.model.ReservedInstancesModificationResult> getModificationResults()
func (jbobject *ServicesEc2ModelReservedInstancesModification) GetModificationResults() []*ServicesEc2ModelReservedInstancesModificationResult {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getModificationResults", "java/util/List")
	if err != nil {
		panic(err)
	}
	retconv := javabind.NewJavaToGoList(javabind.NewJavaToGoCallable())
	dst := new([]*ServicesEc2ModelReservedInstancesModificationResult)
	retconv.Dest(dst)
	if err := retconv.Convert(javabind.ObjectRef(jret)); err != nil {
		panic(err)
	}
	retconv.CleanUp()
	return *dst
}

// public void setModificationResults(java.util.Collection<com.amazonaws.services.ec2.model.ReservedInstancesModificationResult>)
func (jbobject *ServicesEc2ModelReservedInstancesModification) SetModificationResults(a []*ServicesEc2ModelReservedInstancesModificationResult)  {
	conv_a := javabind.NewGoToJavaCollection(javabind.NewGoToJavaCallable())
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "setModificationResults", javabind.Void, conv_a.Value().Cast("java/util/Collection"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()

}

// public com.amazonaws.services.ec2.model.ReservedInstancesModification withModificationResults(com.amazonaws.services.ec2.model.ReservedInstancesModificationResult...)
func (jbobject *ServicesEc2ModelReservedInstancesModification) WithModificationResults(a ...*ServicesEc2ModelReservedInstancesModificationResult) *ServicesEc2ModelReservedInstancesModification {
	conv_a := javabind.NewGoToJavaObjectArray(javabind.NewGoToJavaCallable(), "com/amazonaws/services/ec2/model/ReservedInstancesModificationResult")
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "withModificationResults", "com/amazonaws/services/ec2/model/ReservedInstancesModification", conv_a.Value().Cast("com/amazonaws/services/ec2/model/ReservedInstancesModificationResult"))
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
	unique_x := &ServicesEc2ModelReservedInstancesModification{}
	unique_x.Callable = dst
	return unique_x
}

// public com.amazonaws.services.ec2.model.ReservedInstancesModification withModificationResults(java.util.Collection<com.amazonaws.services.ec2.model.ReservedInstancesModificationResult>)
func (jbobject *ServicesEc2ModelReservedInstancesModification) WithModificationResults2(a []*ServicesEc2ModelReservedInstancesModificationResult) *ServicesEc2ModelReservedInstancesModification {
	conv_a := javabind.NewGoToJavaCollection(javabind.NewGoToJavaCallable())
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "withModificationResults", "com/amazonaws/services/ec2/model/ReservedInstancesModification", conv_a.Value().Cast("java/util/Collection"))
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
	unique_x := &ServicesEc2ModelReservedInstancesModification{}
	unique_x.Callable = dst
	return unique_x
}

// public void setCreateDate(java.util.Date)
func (jbobject *ServicesEc2ModelReservedInstancesModification) SetCreateDate(a time.Time)  {
	conv_a := javabind.NewGoToJavaDate()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "setCreateDate", javabind.Void, conv_a.Value().Cast("java/util/Date"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()

}

// public java.util.Date getCreateDate()
func (jbobject *ServicesEc2ModelReservedInstancesModification) GetCreateDate() time.Time {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getCreateDate", "java/util/Date")
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

// public com.amazonaws.services.ec2.model.ReservedInstancesModification withCreateDate(java.util.Date)
func (jbobject *ServicesEc2ModelReservedInstancesModification) WithCreateDate(a time.Time) *ServicesEc2ModelReservedInstancesModification {
	conv_a := javabind.NewGoToJavaDate()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "withCreateDate", "com/amazonaws/services/ec2/model/ReservedInstancesModification", conv_a.Value().Cast("java/util/Date"))
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
	unique_x := &ServicesEc2ModelReservedInstancesModification{}
	unique_x.Callable = dst
	return unique_x
}

// public void setUpdateDate(java.util.Date)
func (jbobject *ServicesEc2ModelReservedInstancesModification) SetUpdateDate(a time.Time)  {
	conv_a := javabind.NewGoToJavaDate()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "setUpdateDate", javabind.Void, conv_a.Value().Cast("java/util/Date"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()

}

// public java.util.Date getUpdateDate()
func (jbobject *ServicesEc2ModelReservedInstancesModification) GetUpdateDate() time.Time {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getUpdateDate", "java/util/Date")
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

// public com.amazonaws.services.ec2.model.ReservedInstancesModification withUpdateDate(java.util.Date)
func (jbobject *ServicesEc2ModelReservedInstancesModification) WithUpdateDate(a time.Time) *ServicesEc2ModelReservedInstancesModification {
	conv_a := javabind.NewGoToJavaDate()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "withUpdateDate", "com/amazonaws/services/ec2/model/ReservedInstancesModification", conv_a.Value().Cast("java/util/Date"))
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
	unique_x := &ServicesEc2ModelReservedInstancesModification{}
	unique_x.Callable = dst
	return unique_x
}

// public void setEffectiveDate(java.util.Date)
func (jbobject *ServicesEc2ModelReservedInstancesModification) SetEffectiveDate(a time.Time)  {
	conv_a := javabind.NewGoToJavaDate()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "setEffectiveDate", javabind.Void, conv_a.Value().Cast("java/util/Date"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()

}

// public java.util.Date getEffectiveDate()
func (jbobject *ServicesEc2ModelReservedInstancesModification) GetEffectiveDate() time.Time {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getEffectiveDate", "java/util/Date")
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

// public com.amazonaws.services.ec2.model.ReservedInstancesModification withEffectiveDate(java.util.Date)
func (jbobject *ServicesEc2ModelReservedInstancesModification) WithEffectiveDate(a time.Time) *ServicesEc2ModelReservedInstancesModification {
	conv_a := javabind.NewGoToJavaDate()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "withEffectiveDate", "com/amazonaws/services/ec2/model/ReservedInstancesModification", conv_a.Value().Cast("java/util/Date"))
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
	unique_x := &ServicesEc2ModelReservedInstancesModification{}
	unique_x.Callable = dst
	return unique_x
}

// public void setStatus(java.lang.String)
func (jbobject *ServicesEc2ModelReservedInstancesModification) SetStatus(a string)  {
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
func (jbobject *ServicesEc2ModelReservedInstancesModification) GetStatus() string {
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

// public com.amazonaws.services.ec2.model.ReservedInstancesModification withStatus(java.lang.String)
func (jbobject *ServicesEc2ModelReservedInstancesModification) WithStatus(a string) *ServicesEc2ModelReservedInstancesModification {
	conv_a := javabind.NewGoToJavaString()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "withStatus", "com/amazonaws/services/ec2/model/ReservedInstancesModification", conv_a.Value().Cast("java/lang/String"))
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
	unique_x := &ServicesEc2ModelReservedInstancesModification{}
	unique_x.Callable = dst
	return unique_x
}

// public void setStatusMessage(java.lang.String)
func (jbobject *ServicesEc2ModelReservedInstancesModification) SetStatusMessage(a string)  {
	conv_a := javabind.NewGoToJavaString()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "setStatusMessage", javabind.Void, conv_a.Value().Cast("java/lang/String"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()

}

// public java.lang.String getStatusMessage()
func (jbobject *ServicesEc2ModelReservedInstancesModification) GetStatusMessage() string {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getStatusMessage", "java/lang/String")
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

// public com.amazonaws.services.ec2.model.ReservedInstancesModification withStatusMessage(java.lang.String)
func (jbobject *ServicesEc2ModelReservedInstancesModification) WithStatusMessage(a string) *ServicesEc2ModelReservedInstancesModification {
	conv_a := javabind.NewGoToJavaString()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "withStatusMessage", "com/amazonaws/services/ec2/model/ReservedInstancesModification", conv_a.Value().Cast("java/lang/String"))
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
	unique_x := &ServicesEc2ModelReservedInstancesModification{}
	unique_x.Callable = dst
	return unique_x
}

// public void setClientToken(java.lang.String)
func (jbobject *ServicesEc2ModelReservedInstancesModification) SetClientToken(a string)  {
	conv_a := javabind.NewGoToJavaString()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "setClientToken", javabind.Void, conv_a.Value().Cast("java/lang/String"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()

}

// public java.lang.String getClientToken()
func (jbobject *ServicesEc2ModelReservedInstancesModification) GetClientToken() string {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getClientToken", "java/lang/String")
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

// public com.amazonaws.services.ec2.model.ReservedInstancesModification withClientToken(java.lang.String)
func (jbobject *ServicesEc2ModelReservedInstancesModification) WithClientToken(a string) *ServicesEc2ModelReservedInstancesModification {
	conv_a := javabind.NewGoToJavaString()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "withClientToken", "com/amazonaws/services/ec2/model/ReservedInstancesModification", conv_a.Value().Cast("java/lang/String"))
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
	unique_x := &ServicesEc2ModelReservedInstancesModification{}
	unique_x.Callable = dst
	return unique_x
}

// public java.lang.String toString()
func (jbobject *ServicesEc2ModelReservedInstancesModification) ToString() string {
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
func (jbobject *ServicesEc2ModelReservedInstancesModification) Equals(a interface{}) bool {
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
func (jbobject *ServicesEc2ModelReservedInstancesModification) HashCode() int {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "hashCode", javabind.Int)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

// public com.amazonaws.services.ec2.model.ReservedInstancesModification clone()
func (jbobject *ServicesEc2ModelReservedInstancesModification) Clone() *ServicesEc2ModelReservedInstancesModification {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "clone", "com/amazonaws/services/ec2/model/ReservedInstancesModification")
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
	unique_x := &ServicesEc2ModelReservedInstancesModification{}
	unique_x.Callable = dst
	return unique_x
}

// public java.lang.Object clone() throws java.lang.CloneNotSupportedException
func (jbobject *ServicesEc2ModelReservedInstancesModification) Clone2() (*JavaLangObject, error) {
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


