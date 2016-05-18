package ec2

import "time"
import "github.com/timob/javabind"

type ServicesEc2ModelSpotFleetRequestConfigInterface interface {
	JavaLangObjectInterface

	// public void setSpotFleetRequestId(java.lang.String)
	SetSpotFleetRequestId(a string) 

	// public java.lang.String getSpotFleetRequestId()
	GetSpotFleetRequestId() string

	// public com.amazonaws.services.ec2.model.SpotFleetRequestConfig withSpotFleetRequestId(java.lang.String)
	WithSpotFleetRequestId(a string) *ServicesEc2ModelSpotFleetRequestConfig

	// public void setSpotFleetRequestState(java.lang.String)
	SetSpotFleetRequestState2(a string) 

	// public java.lang.String getSpotFleetRequestState()
	GetSpotFleetRequestState() string

	// public com.amazonaws.services.ec2.model.SpotFleetRequestConfig withSpotFleetRequestState(java.lang.String)
	WithSpotFleetRequestState2(a string) *ServicesEc2ModelSpotFleetRequestConfig

	// public void setSpotFleetRequestState(com.amazonaws.services.ec2.model.BatchState)
	SetSpotFleetRequestState(a ServicesEc2ModelBatchStateInterface) 

	// public com.amazonaws.services.ec2.model.SpotFleetRequestConfig withSpotFleetRequestState(com.amazonaws.services.ec2.model.BatchState)
	WithSpotFleetRequestState(a ServicesEc2ModelBatchStateInterface) *ServicesEc2ModelSpotFleetRequestConfig

	// public void setSpotFleetRequestConfig(com.amazonaws.services.ec2.model.SpotFleetRequestConfigData)
	SetSpotFleetRequestConfig(a ServicesEc2ModelSpotFleetRequestConfigDataInterface) 

	// public com.amazonaws.services.ec2.model.SpotFleetRequestConfigData getSpotFleetRequestConfig()
	GetSpotFleetRequestConfig() *ServicesEc2ModelSpotFleetRequestConfigData

	// public com.amazonaws.services.ec2.model.SpotFleetRequestConfig withSpotFleetRequestConfig(com.amazonaws.services.ec2.model.SpotFleetRequestConfigData)
	WithSpotFleetRequestConfig(a ServicesEc2ModelSpotFleetRequestConfigDataInterface) *ServicesEc2ModelSpotFleetRequestConfig

	// public void setCreateTime(java.util.Date)
	SetCreateTime(a time.Time) 

	// public java.util.Date getCreateTime()
	GetCreateTime() time.Time

	// public com.amazonaws.services.ec2.model.SpotFleetRequestConfig withCreateTime(java.util.Date)
	WithCreateTime(a time.Time) *ServicesEc2ModelSpotFleetRequestConfig

	// public com.amazonaws.services.ec2.model.SpotFleetRequestConfig clone()
	Clone() *ServicesEc2ModelSpotFleetRequestConfig
}

type ServicesEc2ModelSpotFleetRequestConfig struct {
	JavaLangObject
}

// public com.amazonaws.services.ec2.model.SpotFleetRequestConfig()
func NewServicesEc2ModelSpotFleetRequestConfig() (*ServicesEc2ModelSpotFleetRequestConfig) {

	obj, err := javabind.GetEnv().NewObject("com/amazonaws/services/ec2/model/SpotFleetRequestConfig")
	if err != nil {
		panic(err)
	}
	x := &ServicesEc2ModelSpotFleetRequestConfig{}
	x.Callable = &javabind.Callable{obj}
	return x
}

// public void setSpotFleetRequestId(java.lang.String)
func (jbobject *ServicesEc2ModelSpotFleetRequestConfig) SetSpotFleetRequestId(a string)  {
	conv_a := javabind.NewGoToJavaString()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "setSpotFleetRequestId", javabind.Void, conv_a.Value().Cast("java/lang/String"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()

}

// public java.lang.String getSpotFleetRequestId()
func (jbobject *ServicesEc2ModelSpotFleetRequestConfig) GetSpotFleetRequestId() string {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getSpotFleetRequestId", "java/lang/String")
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

// public com.amazonaws.services.ec2.model.SpotFleetRequestConfig withSpotFleetRequestId(java.lang.String)
func (jbobject *ServicesEc2ModelSpotFleetRequestConfig) WithSpotFleetRequestId(a string) *ServicesEc2ModelSpotFleetRequestConfig {
	conv_a := javabind.NewGoToJavaString()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "withSpotFleetRequestId", "com/amazonaws/services/ec2/model/SpotFleetRequestConfig", conv_a.Value().Cast("java/lang/String"))
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
	unique_x := &ServicesEc2ModelSpotFleetRequestConfig{}
	unique_x.Callable = dst
	return unique_x
}

// public void setSpotFleetRequestState(java.lang.String)
func (jbobject *ServicesEc2ModelSpotFleetRequestConfig) SetSpotFleetRequestState2(a string)  {
	conv_a := javabind.NewGoToJavaString()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "setSpotFleetRequestState", javabind.Void, conv_a.Value().Cast("java/lang/String"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()

}

// public java.lang.String getSpotFleetRequestState()
func (jbobject *ServicesEc2ModelSpotFleetRequestConfig) GetSpotFleetRequestState() string {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getSpotFleetRequestState", "java/lang/String")
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

// public com.amazonaws.services.ec2.model.SpotFleetRequestConfig withSpotFleetRequestState(java.lang.String)
func (jbobject *ServicesEc2ModelSpotFleetRequestConfig) WithSpotFleetRequestState2(a string) *ServicesEc2ModelSpotFleetRequestConfig {
	conv_a := javabind.NewGoToJavaString()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "withSpotFleetRequestState", "com/amazonaws/services/ec2/model/SpotFleetRequestConfig", conv_a.Value().Cast("java/lang/String"))
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
	unique_x := &ServicesEc2ModelSpotFleetRequestConfig{}
	unique_x.Callable = dst
	return unique_x
}

// public void setSpotFleetRequestState(com.amazonaws.services.ec2.model.BatchState)
func (jbobject *ServicesEc2ModelSpotFleetRequestConfig) SetSpotFleetRequestState(a ServicesEc2ModelBatchStateInterface)  {
	conv_a := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "setSpotFleetRequestState", javabind.Void, conv_a.Value().Cast("com/amazonaws/services/ec2/model/BatchState"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()

}

// public com.amazonaws.services.ec2.model.SpotFleetRequestConfig withSpotFleetRequestState(com.amazonaws.services.ec2.model.BatchState)
func (jbobject *ServicesEc2ModelSpotFleetRequestConfig) WithSpotFleetRequestState(a ServicesEc2ModelBatchStateInterface) *ServicesEc2ModelSpotFleetRequestConfig {
	conv_a := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "withSpotFleetRequestState", "com/amazonaws/services/ec2/model/SpotFleetRequestConfig", conv_a.Value().Cast("com/amazonaws/services/ec2/model/BatchState"))
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
	unique_x := &ServicesEc2ModelSpotFleetRequestConfig{}
	unique_x.Callable = dst
	return unique_x
}

// public void setSpotFleetRequestConfig(com.amazonaws.services.ec2.model.SpotFleetRequestConfigData)
func (jbobject *ServicesEc2ModelSpotFleetRequestConfig) SetSpotFleetRequestConfig(a ServicesEc2ModelSpotFleetRequestConfigDataInterface)  {
	conv_a := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "setSpotFleetRequestConfig", javabind.Void, conv_a.Value().Cast("com/amazonaws/services/ec2/model/SpotFleetRequestConfigData"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()

}

// public com.amazonaws.services.ec2.model.SpotFleetRequestConfigData getSpotFleetRequestConfig()
func (jbobject *ServicesEc2ModelSpotFleetRequestConfig) GetSpotFleetRequestConfig() *ServicesEc2ModelSpotFleetRequestConfigData {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getSpotFleetRequestConfig", "com/amazonaws/services/ec2/model/SpotFleetRequestConfigData")
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
	unique_x := &ServicesEc2ModelSpotFleetRequestConfigData{}
	unique_x.Callable = dst
	return unique_x
}

// public com.amazonaws.services.ec2.model.SpotFleetRequestConfig withSpotFleetRequestConfig(com.amazonaws.services.ec2.model.SpotFleetRequestConfigData)
func (jbobject *ServicesEc2ModelSpotFleetRequestConfig) WithSpotFleetRequestConfig(a ServicesEc2ModelSpotFleetRequestConfigDataInterface) *ServicesEc2ModelSpotFleetRequestConfig {
	conv_a := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "withSpotFleetRequestConfig", "com/amazonaws/services/ec2/model/SpotFleetRequestConfig", conv_a.Value().Cast("com/amazonaws/services/ec2/model/SpotFleetRequestConfigData"))
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
	unique_x := &ServicesEc2ModelSpotFleetRequestConfig{}
	unique_x.Callable = dst
	return unique_x
}

// public void setCreateTime(java.util.Date)
func (jbobject *ServicesEc2ModelSpotFleetRequestConfig) SetCreateTime(a time.Time)  {
	conv_a := javabind.NewGoToJavaDate()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "setCreateTime", javabind.Void, conv_a.Value().Cast("java/util/Date"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()

}

// public java.util.Date getCreateTime()
func (jbobject *ServicesEc2ModelSpotFleetRequestConfig) GetCreateTime() time.Time {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getCreateTime", "java/util/Date")
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

// public com.amazonaws.services.ec2.model.SpotFleetRequestConfig withCreateTime(java.util.Date)
func (jbobject *ServicesEc2ModelSpotFleetRequestConfig) WithCreateTime(a time.Time) *ServicesEc2ModelSpotFleetRequestConfig {
	conv_a := javabind.NewGoToJavaDate()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "withCreateTime", "com/amazonaws/services/ec2/model/SpotFleetRequestConfig", conv_a.Value().Cast("java/util/Date"))
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
	unique_x := &ServicesEc2ModelSpotFleetRequestConfig{}
	unique_x.Callable = dst
	return unique_x
}

// public java.lang.String toString()
func (jbobject *ServicesEc2ModelSpotFleetRequestConfig) ToString() string {
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
func (jbobject *ServicesEc2ModelSpotFleetRequestConfig) Equals(a interface{}) bool {
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
func (jbobject *ServicesEc2ModelSpotFleetRequestConfig) HashCode() int {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "hashCode", javabind.Int)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

// public com.amazonaws.services.ec2.model.SpotFleetRequestConfig clone()
func (jbobject *ServicesEc2ModelSpotFleetRequestConfig) Clone() *ServicesEc2ModelSpotFleetRequestConfig {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "clone", "com/amazonaws/services/ec2/model/SpotFleetRequestConfig")
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
	unique_x := &ServicesEc2ModelSpotFleetRequestConfig{}
	unique_x.Callable = dst
	return unique_x
}

// public java.lang.Object clone() throws java.lang.CloneNotSupportedException
func (jbobject *ServicesEc2ModelSpotFleetRequestConfig) Clone2() (*JavaLangObject, error) {
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


