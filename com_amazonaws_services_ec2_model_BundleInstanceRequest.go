package ec2

import "github.com/timob/javabind"

type ServicesEc2ModelBundleInstanceRequestInterface interface {
	AmazonWebServiceRequestInterface

	// public void setInstanceId(java.lang.String)
	SetInstanceId(a string) 

	// public java.lang.String getInstanceId()
	GetInstanceId() string

	// public com.amazonaws.services.ec2.model.BundleInstanceRequest withInstanceId(java.lang.String)
	WithInstanceId(a string) *ServicesEc2ModelBundleInstanceRequest

	// public void setStorage(com.amazonaws.services.ec2.model.Storage)
	SetStorage(a ServicesEc2ModelStorageInterface) 

	// public com.amazonaws.services.ec2.model.Storage getStorage()
	GetStorage() *ServicesEc2ModelStorage

	// public com.amazonaws.services.ec2.model.BundleInstanceRequest withStorage(com.amazonaws.services.ec2.model.Storage)
	WithStorage(a ServicesEc2ModelStorageInterface) *ServicesEc2ModelBundleInstanceRequest

	// public com.amazonaws.Request<com.amazonaws.services.ec2.model.BundleInstanceRequest> getDryRunRequest()
	GetDryRunRequest() *Request

	// public com.amazonaws.services.ec2.model.BundleInstanceRequest clone()
	Clone3() *ServicesEc2ModelBundleInstanceRequest
}

type ServicesEc2ModelBundleInstanceRequest struct {
	AmazonWebServiceRequest
}

// public com.amazonaws.services.ec2.model.BundleInstanceRequest()
func NewServicesEc2ModelBundleInstanceRequest() (*ServicesEc2ModelBundleInstanceRequest) {

	obj, err := javabind.GetEnv().NewObject("com/amazonaws/services/ec2/model/BundleInstanceRequest")
	if err != nil {
		panic(err)
	}
	x := &ServicesEc2ModelBundleInstanceRequest{}
	x.Callable = &javabind.Callable{obj}
	return x
}

// public com.amazonaws.services.ec2.model.BundleInstanceRequest(java.lang.String, com.amazonaws.services.ec2.model.Storage)
func NewServicesEc2ModelBundleInstanceRequest2(a string, b ServicesEc2ModelStorageInterface) (*ServicesEc2ModelBundleInstanceRequest) {
	conv_a := javabind.NewGoToJavaString()
	conv_b := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	if err := conv_b.Convert(b); err != nil {
		panic(err)
	}

	obj, err := javabind.GetEnv().NewObject("com/amazonaws/services/ec2/model/BundleInstanceRequest", conv_a.Value().Cast("java/lang/String"), conv_b.Value().Cast("com/amazonaws/services/ec2/model/Storage"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()
	conv_b.CleanUp()
	x := &ServicesEc2ModelBundleInstanceRequest{}
	x.Callable = &javabind.Callable{obj}
	return x
}

// public void setInstanceId(java.lang.String)
func (jbobject *ServicesEc2ModelBundleInstanceRequest) SetInstanceId(a string)  {
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
func (jbobject *ServicesEc2ModelBundleInstanceRequest) GetInstanceId() string {
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

// public com.amazonaws.services.ec2.model.BundleInstanceRequest withInstanceId(java.lang.String)
func (jbobject *ServicesEc2ModelBundleInstanceRequest) WithInstanceId(a string) *ServicesEc2ModelBundleInstanceRequest {
	conv_a := javabind.NewGoToJavaString()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "withInstanceId", "com/amazonaws/services/ec2/model/BundleInstanceRequest", conv_a.Value().Cast("java/lang/String"))
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
	unique_x := &ServicesEc2ModelBundleInstanceRequest{}
	unique_x.Callable = dst
	return unique_x
}

// public void setStorage(com.amazonaws.services.ec2.model.Storage)
func (jbobject *ServicesEc2ModelBundleInstanceRequest) SetStorage(a ServicesEc2ModelStorageInterface)  {
	conv_a := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "setStorage", javabind.Void, conv_a.Value().Cast("com/amazonaws/services/ec2/model/Storage"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()

}

// public com.amazonaws.services.ec2.model.Storage getStorage()
func (jbobject *ServicesEc2ModelBundleInstanceRequest) GetStorage() *ServicesEc2ModelStorage {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getStorage", "com/amazonaws/services/ec2/model/Storage")
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
	unique_x := &ServicesEc2ModelStorage{}
	unique_x.Callable = dst
	return unique_x
}

// public com.amazonaws.services.ec2.model.BundleInstanceRequest withStorage(com.amazonaws.services.ec2.model.Storage)
func (jbobject *ServicesEc2ModelBundleInstanceRequest) WithStorage(a ServicesEc2ModelStorageInterface) *ServicesEc2ModelBundleInstanceRequest {
	conv_a := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "withStorage", "com/amazonaws/services/ec2/model/BundleInstanceRequest", conv_a.Value().Cast("com/amazonaws/services/ec2/model/Storage"))
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
	unique_x := &ServicesEc2ModelBundleInstanceRequest{}
	unique_x.Callable = dst
	return unique_x
}

// public com.amazonaws.Request<com.amazonaws.services.ec2.model.BundleInstanceRequest> getDryRunRequest()
func (jbobject *ServicesEc2ModelBundleInstanceRequest) GetDryRunRequest() *Request {
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
func (jbobject *ServicesEc2ModelBundleInstanceRequest) ToString() string {
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
func (jbobject *ServicesEc2ModelBundleInstanceRequest) Equals(a interface{}) bool {
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
func (jbobject *ServicesEc2ModelBundleInstanceRequest) HashCode() int {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "hashCode", javabind.Int)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

// public com.amazonaws.services.ec2.model.BundleInstanceRequest clone()
func (jbobject *ServicesEc2ModelBundleInstanceRequest) Clone3() *ServicesEc2ModelBundleInstanceRequest {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "clone", "com/amazonaws/services/ec2/model/BundleInstanceRequest")
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
	unique_x := &ServicesEc2ModelBundleInstanceRequest{}
	unique_x.Callable = dst
	return unique_x
}

// public com.amazonaws.AmazonWebServiceRequest clone()
func (jbobject *ServicesEc2ModelBundleInstanceRequest) Clone() *AmazonWebServiceRequest {
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
func (jbobject *ServicesEc2ModelBundleInstanceRequest) Clone2() (*JavaLangObject, error) {
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


