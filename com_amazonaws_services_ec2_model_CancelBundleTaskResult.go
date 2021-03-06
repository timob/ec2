package ec2

import "github.com/timob/javabind"

type ServicesEc2ModelCancelBundleTaskResultInterface interface {
	JavaLangObjectInterface

	// public void setBundleTask(com.amazonaws.services.ec2.model.BundleTask)
	SetBundleTask(a ServicesEc2ModelBundleTaskInterface) 

	// public com.amazonaws.services.ec2.model.BundleTask getBundleTask()
	GetBundleTask() *ServicesEc2ModelBundleTask

	// public com.amazonaws.services.ec2.model.CancelBundleTaskResult withBundleTask(com.amazonaws.services.ec2.model.BundleTask)
	WithBundleTask(a ServicesEc2ModelBundleTaskInterface) *ServicesEc2ModelCancelBundleTaskResult

	// public com.amazonaws.services.ec2.model.CancelBundleTaskResult clone()
	Clone() *ServicesEc2ModelCancelBundleTaskResult
}

type ServicesEc2ModelCancelBundleTaskResult struct {
	JavaLangObject
}

// public com.amazonaws.services.ec2.model.CancelBundleTaskResult()
func NewServicesEc2ModelCancelBundleTaskResult() (*ServicesEc2ModelCancelBundleTaskResult) {

	obj, err := javabind.GetEnv().NewObject("com/amazonaws/services/ec2/model/CancelBundleTaskResult")
	if err != nil {
		panic(err)
	}
	x := &ServicesEc2ModelCancelBundleTaskResult{}
	x.Callable = &javabind.Callable{obj}
	return x
}

// public void setBundleTask(com.amazonaws.services.ec2.model.BundleTask)
func (jbobject *ServicesEc2ModelCancelBundleTaskResult) SetBundleTask(a ServicesEc2ModelBundleTaskInterface)  {
	conv_a := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "setBundleTask", javabind.Void, conv_a.Value().Cast("com/amazonaws/services/ec2/model/BundleTask"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()

}

// public com.amazonaws.services.ec2.model.BundleTask getBundleTask()
func (jbobject *ServicesEc2ModelCancelBundleTaskResult) GetBundleTask() *ServicesEc2ModelBundleTask {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getBundleTask", "com/amazonaws/services/ec2/model/BundleTask")
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
	unique_x := &ServicesEc2ModelBundleTask{}
	unique_x.Callable = dst
	return unique_x
}

// public com.amazonaws.services.ec2.model.CancelBundleTaskResult withBundleTask(com.amazonaws.services.ec2.model.BundleTask)
func (jbobject *ServicesEc2ModelCancelBundleTaskResult) WithBundleTask(a ServicesEc2ModelBundleTaskInterface) *ServicesEc2ModelCancelBundleTaskResult {
	conv_a := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "withBundleTask", "com/amazonaws/services/ec2/model/CancelBundleTaskResult", conv_a.Value().Cast("com/amazonaws/services/ec2/model/BundleTask"))
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
	unique_x := &ServicesEc2ModelCancelBundleTaskResult{}
	unique_x.Callable = dst
	return unique_x
}

// public java.lang.String toString()
func (jbobject *ServicesEc2ModelCancelBundleTaskResult) ToString() string {
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
func (jbobject *ServicesEc2ModelCancelBundleTaskResult) Equals(a interface{}) bool {
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
func (jbobject *ServicesEc2ModelCancelBundleTaskResult) HashCode() int {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "hashCode", javabind.Int)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

// public com.amazonaws.services.ec2.model.CancelBundleTaskResult clone()
func (jbobject *ServicesEc2ModelCancelBundleTaskResult) Clone() *ServicesEc2ModelCancelBundleTaskResult {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "clone", "com/amazonaws/services/ec2/model/CancelBundleTaskResult")
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
	unique_x := &ServicesEc2ModelCancelBundleTaskResult{}
	unique_x.Callable = dst
	return unique_x
}

// public java.lang.Object clone() throws java.lang.CloneNotSupportedException
func (jbobject *ServicesEc2ModelCancelBundleTaskResult) Clone2() (*JavaLangObject, error) {
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


