package ec2

import "github.com/timob/javabind"

type ServicesEc2ModelReservedInstancesModificationResultInterface interface {
	JavaLangObjectInterface

	// public void setReservedInstancesId(java.lang.String)
	SetReservedInstancesId(a string) 

	// public java.lang.String getReservedInstancesId()
	GetReservedInstancesId() string

	// public com.amazonaws.services.ec2.model.ReservedInstancesModificationResult withReservedInstancesId(java.lang.String)
	WithReservedInstancesId(a string) *ServicesEc2ModelReservedInstancesModificationResult

	// public void setTargetConfiguration(com.amazonaws.services.ec2.model.ReservedInstancesConfiguration)
	SetTargetConfiguration(a ServicesEc2ModelReservedInstancesConfigurationInterface) 

	// public com.amazonaws.services.ec2.model.ReservedInstancesConfiguration getTargetConfiguration()
	GetTargetConfiguration() *ServicesEc2ModelReservedInstancesConfiguration

	// public com.amazonaws.services.ec2.model.ReservedInstancesModificationResult withTargetConfiguration(com.amazonaws.services.ec2.model.ReservedInstancesConfiguration)
	WithTargetConfiguration(a ServicesEc2ModelReservedInstancesConfigurationInterface) *ServicesEc2ModelReservedInstancesModificationResult

	// public com.amazonaws.services.ec2.model.ReservedInstancesModificationResult clone()
	Clone() *ServicesEc2ModelReservedInstancesModificationResult
}

type ServicesEc2ModelReservedInstancesModificationResult struct {
	JavaLangObject
}

// public com.amazonaws.services.ec2.model.ReservedInstancesModificationResult()
func NewServicesEc2ModelReservedInstancesModificationResult() (*ServicesEc2ModelReservedInstancesModificationResult) {

	obj, err := javabind.GetEnv().NewObject("com/amazonaws/services/ec2/model/ReservedInstancesModificationResult")
	if err != nil {
		panic(err)
	}
	x := &ServicesEc2ModelReservedInstancesModificationResult{}
	x.Callable = &javabind.Callable{obj}
	return x
}

// public void setReservedInstancesId(java.lang.String)
func (jbobject *ServicesEc2ModelReservedInstancesModificationResult) SetReservedInstancesId(a string)  {
	conv_a := javabind.NewGoToJavaString()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "setReservedInstancesId", javabind.Void, conv_a.Value().Cast("java/lang/String"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()

}

// public java.lang.String getReservedInstancesId()
func (jbobject *ServicesEc2ModelReservedInstancesModificationResult) GetReservedInstancesId() string {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getReservedInstancesId", "java/lang/String")
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

// public com.amazonaws.services.ec2.model.ReservedInstancesModificationResult withReservedInstancesId(java.lang.String)
func (jbobject *ServicesEc2ModelReservedInstancesModificationResult) WithReservedInstancesId(a string) *ServicesEc2ModelReservedInstancesModificationResult {
	conv_a := javabind.NewGoToJavaString()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "withReservedInstancesId", "com/amazonaws/services/ec2/model/ReservedInstancesModificationResult", conv_a.Value().Cast("java/lang/String"))
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
	unique_x := &ServicesEc2ModelReservedInstancesModificationResult{}
	unique_x.Callable = dst
	return unique_x
}

// public void setTargetConfiguration(com.amazonaws.services.ec2.model.ReservedInstancesConfiguration)
func (jbobject *ServicesEc2ModelReservedInstancesModificationResult) SetTargetConfiguration(a ServicesEc2ModelReservedInstancesConfigurationInterface)  {
	conv_a := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "setTargetConfiguration", javabind.Void, conv_a.Value().Cast("com/amazonaws/services/ec2/model/ReservedInstancesConfiguration"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()

}

// public com.amazonaws.services.ec2.model.ReservedInstancesConfiguration getTargetConfiguration()
func (jbobject *ServicesEc2ModelReservedInstancesModificationResult) GetTargetConfiguration() *ServicesEc2ModelReservedInstancesConfiguration {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getTargetConfiguration", "com/amazonaws/services/ec2/model/ReservedInstancesConfiguration")
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
	unique_x := &ServicesEc2ModelReservedInstancesConfiguration{}
	unique_x.Callable = dst
	return unique_x
}

// public com.amazonaws.services.ec2.model.ReservedInstancesModificationResult withTargetConfiguration(com.amazonaws.services.ec2.model.ReservedInstancesConfiguration)
func (jbobject *ServicesEc2ModelReservedInstancesModificationResult) WithTargetConfiguration(a ServicesEc2ModelReservedInstancesConfigurationInterface) *ServicesEc2ModelReservedInstancesModificationResult {
	conv_a := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "withTargetConfiguration", "com/amazonaws/services/ec2/model/ReservedInstancesModificationResult", conv_a.Value().Cast("com/amazonaws/services/ec2/model/ReservedInstancesConfiguration"))
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
	unique_x := &ServicesEc2ModelReservedInstancesModificationResult{}
	unique_x.Callable = dst
	return unique_x
}

// public java.lang.String toString()
func (jbobject *ServicesEc2ModelReservedInstancesModificationResult) ToString() string {
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
func (jbobject *ServicesEc2ModelReservedInstancesModificationResult) Equals(a interface{}) bool {
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
func (jbobject *ServicesEc2ModelReservedInstancesModificationResult) HashCode() int {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "hashCode", javabind.Int)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

// public com.amazonaws.services.ec2.model.ReservedInstancesModificationResult clone()
func (jbobject *ServicesEc2ModelReservedInstancesModificationResult) Clone() *ServicesEc2ModelReservedInstancesModificationResult {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "clone", "com/amazonaws/services/ec2/model/ReservedInstancesModificationResult")
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
	unique_x := &ServicesEc2ModelReservedInstancesModificationResult{}
	unique_x.Callable = dst
	return unique_x
}

// public java.lang.Object clone() throws java.lang.CloneNotSupportedException
func (jbobject *ServicesEc2ModelReservedInstancesModificationResult) Clone2() (*JavaLangObject, error) {
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


