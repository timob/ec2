package ec2

import "github.com/timob/javabind"

type ServicesEc2ModelImportVolumeResultInterface interface {
	JavaLangObjectInterface

	// public void setConversionTask(com.amazonaws.services.ec2.model.ConversionTask)
	SetConversionTask(a ServicesEc2ModelConversionTaskInterface) 

	// public com.amazonaws.services.ec2.model.ConversionTask getConversionTask()
	GetConversionTask() *ServicesEc2ModelConversionTask

	// public com.amazonaws.services.ec2.model.ImportVolumeResult withConversionTask(com.amazonaws.services.ec2.model.ConversionTask)
	WithConversionTask(a ServicesEc2ModelConversionTaskInterface) *ServicesEc2ModelImportVolumeResult

	// public com.amazonaws.services.ec2.model.ImportVolumeResult clone()
	Clone() *ServicesEc2ModelImportVolumeResult
}

type ServicesEc2ModelImportVolumeResult struct {
	JavaLangObject
}

// public com.amazonaws.services.ec2.model.ImportVolumeResult()
func NewServicesEc2ModelImportVolumeResult() (*ServicesEc2ModelImportVolumeResult) {

	obj, err := javabind.GetEnv().NewObject("com/amazonaws/services/ec2/model/ImportVolumeResult")
	if err != nil {
		panic(err)
	}
	x := &ServicesEc2ModelImportVolumeResult{}
	x.Callable = &javabind.Callable{obj}
	return x
}

// public void setConversionTask(com.amazonaws.services.ec2.model.ConversionTask)
func (jbobject *ServicesEc2ModelImportVolumeResult) SetConversionTask(a ServicesEc2ModelConversionTaskInterface)  {
	conv_a := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "setConversionTask", javabind.Void, conv_a.Value().Cast("com/amazonaws/services/ec2/model/ConversionTask"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()

}

// public com.amazonaws.services.ec2.model.ConversionTask getConversionTask()
func (jbobject *ServicesEc2ModelImportVolumeResult) GetConversionTask() *ServicesEc2ModelConversionTask {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getConversionTask", "com/amazonaws/services/ec2/model/ConversionTask")
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
	unique_x := &ServicesEc2ModelConversionTask{}
	unique_x.Callable = dst
	return unique_x
}

// public com.amazonaws.services.ec2.model.ImportVolumeResult withConversionTask(com.amazonaws.services.ec2.model.ConversionTask)
func (jbobject *ServicesEc2ModelImportVolumeResult) WithConversionTask(a ServicesEc2ModelConversionTaskInterface) *ServicesEc2ModelImportVolumeResult {
	conv_a := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "withConversionTask", "com/amazonaws/services/ec2/model/ImportVolumeResult", conv_a.Value().Cast("com/amazonaws/services/ec2/model/ConversionTask"))
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
	unique_x := &ServicesEc2ModelImportVolumeResult{}
	unique_x.Callable = dst
	return unique_x
}

// public java.lang.String toString()
func (jbobject *ServicesEc2ModelImportVolumeResult) ToString() string {
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
func (jbobject *ServicesEc2ModelImportVolumeResult) Equals(a interface{}) bool {
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
func (jbobject *ServicesEc2ModelImportVolumeResult) HashCode() int {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "hashCode", javabind.Int)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

// public com.amazonaws.services.ec2.model.ImportVolumeResult clone()
func (jbobject *ServicesEc2ModelImportVolumeResult) Clone() *ServicesEc2ModelImportVolumeResult {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "clone", "com/amazonaws/services/ec2/model/ImportVolumeResult")
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
	unique_x := &ServicesEc2ModelImportVolumeResult{}
	unique_x.Callable = dst
	return unique_x
}

// public java.lang.Object clone() throws java.lang.CloneNotSupportedException
func (jbobject *ServicesEc2ModelImportVolumeResult) Clone2() (*JavaLangObject, error) {
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


