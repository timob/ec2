package ec2

import "github.com/timob/javabind"

type ServicesEc2ModelInstanceExportDetailsInterface interface {
	JavaLangObjectInterface

	// public void setInstanceId(java.lang.String)
	SetInstanceId(a string) 

	// public java.lang.String getInstanceId()
	GetInstanceId() string

	// public com.amazonaws.services.ec2.model.InstanceExportDetails withInstanceId(java.lang.String)
	WithInstanceId(a string) *ServicesEc2ModelInstanceExportDetails

	// public void setTargetEnvironment(java.lang.String)
	SetTargetEnvironment2(a string) 

	// public java.lang.String getTargetEnvironment()
	GetTargetEnvironment() string

	// public com.amazonaws.services.ec2.model.InstanceExportDetails withTargetEnvironment(java.lang.String)
	WithTargetEnvironment2(a string) *ServicesEc2ModelInstanceExportDetails

	// public void setTargetEnvironment(com.amazonaws.services.ec2.model.ExportEnvironment)
	SetTargetEnvironment(a ServicesEc2ModelExportEnvironmentInterface) 

	// public com.amazonaws.services.ec2.model.InstanceExportDetails withTargetEnvironment(com.amazonaws.services.ec2.model.ExportEnvironment)
	WithTargetEnvironment(a ServicesEc2ModelExportEnvironmentInterface) *ServicesEc2ModelInstanceExportDetails

	// public com.amazonaws.services.ec2.model.InstanceExportDetails clone()
	Clone() *ServicesEc2ModelInstanceExportDetails
}

type ServicesEc2ModelInstanceExportDetails struct {
	JavaLangObject
}

// public com.amazonaws.services.ec2.model.InstanceExportDetails()
func NewServicesEc2ModelInstanceExportDetails() (*ServicesEc2ModelInstanceExportDetails) {

	obj, err := javabind.GetEnv().NewObject("com/amazonaws/services/ec2/model/InstanceExportDetails")
	if err != nil {
		panic(err)
	}
	x := &ServicesEc2ModelInstanceExportDetails{}
	x.Callable = &javabind.Callable{obj}
	return x
}

// public void setInstanceId(java.lang.String)
func (jbobject *ServicesEc2ModelInstanceExportDetails) SetInstanceId(a string)  {
	conv_a := javabind.NewGoToJavaString()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "setInstanceId", javabind.Void, conv_a.Value().Cast("java/lang/String"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()

}

// public java.lang.String getInstanceId()
func (jbobject *ServicesEc2ModelInstanceExportDetails) GetInstanceId() string {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getInstanceId", "java/lang/String")
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

// public com.amazonaws.services.ec2.model.InstanceExportDetails withInstanceId(java.lang.String)
func (jbobject *ServicesEc2ModelInstanceExportDetails) WithInstanceId(a string) *ServicesEc2ModelInstanceExportDetails {
	conv_a := javabind.NewGoToJavaString()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "withInstanceId", "com/amazonaws/services/ec2/model/InstanceExportDetails", conv_a.Value().Cast("java/lang/String"))
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
	unique_x := &ServicesEc2ModelInstanceExportDetails{}
	unique_x.Callable = dst
	return unique_x
}

// public void setTargetEnvironment(java.lang.String)
func (jbobject *ServicesEc2ModelInstanceExportDetails) SetTargetEnvironment2(a string)  {
	conv_a := javabind.NewGoToJavaString()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "setTargetEnvironment", javabind.Void, conv_a.Value().Cast("java/lang/String"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()

}

// public java.lang.String getTargetEnvironment()
func (jbobject *ServicesEc2ModelInstanceExportDetails) GetTargetEnvironment() string {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getTargetEnvironment", "java/lang/String")
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

// public com.amazonaws.services.ec2.model.InstanceExportDetails withTargetEnvironment(java.lang.String)
func (jbobject *ServicesEc2ModelInstanceExportDetails) WithTargetEnvironment2(a string) *ServicesEc2ModelInstanceExportDetails {
	conv_a := javabind.NewGoToJavaString()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "withTargetEnvironment", "com/amazonaws/services/ec2/model/InstanceExportDetails", conv_a.Value().Cast("java/lang/String"))
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
	unique_x := &ServicesEc2ModelInstanceExportDetails{}
	unique_x.Callable = dst
	return unique_x
}

// public void setTargetEnvironment(com.amazonaws.services.ec2.model.ExportEnvironment)
func (jbobject *ServicesEc2ModelInstanceExportDetails) SetTargetEnvironment(a ServicesEc2ModelExportEnvironmentInterface)  {
	conv_a := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "setTargetEnvironment", javabind.Void, conv_a.Value().Cast("com/amazonaws/services/ec2/model/ExportEnvironment"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()

}

// public com.amazonaws.services.ec2.model.InstanceExportDetails withTargetEnvironment(com.amazonaws.services.ec2.model.ExportEnvironment)
func (jbobject *ServicesEc2ModelInstanceExportDetails) WithTargetEnvironment(a ServicesEc2ModelExportEnvironmentInterface) *ServicesEc2ModelInstanceExportDetails {
	conv_a := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "withTargetEnvironment", "com/amazonaws/services/ec2/model/InstanceExportDetails", conv_a.Value().Cast("com/amazonaws/services/ec2/model/ExportEnvironment"))
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
	unique_x := &ServicesEc2ModelInstanceExportDetails{}
	unique_x.Callable = dst
	return unique_x
}

// public java.lang.String toString()
func (jbobject *ServicesEc2ModelInstanceExportDetails) ToString() string {
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
func (jbobject *ServicesEc2ModelInstanceExportDetails) Equals(a interface{}) bool {
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
func (jbobject *ServicesEc2ModelInstanceExportDetails) HashCode() int {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "hashCode", javabind.Int)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

// public com.amazonaws.services.ec2.model.InstanceExportDetails clone()
func (jbobject *ServicesEc2ModelInstanceExportDetails) Clone() *ServicesEc2ModelInstanceExportDetails {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "clone", "com/amazonaws/services/ec2/model/InstanceExportDetails")
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
	unique_x := &ServicesEc2ModelInstanceExportDetails{}
	unique_x.Callable = dst
	return unique_x
}

// public java.lang.Object clone() throws java.lang.CloneNotSupportedException
func (jbobject *ServicesEc2ModelInstanceExportDetails) Clone2() (*JavaLangObject, error) {
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


