package ec2

import "github.com/timob/javabind"

type ServicesEc2ModelDescribeImagesResultInterface interface {
	JavaLangObjectInterface

	// public java.util.List<com.amazonaws.services.ec2.model.Image> getImages()
	GetImages() []*ServicesEc2ModelImage

	// public void setImages(java.util.Collection<com.amazonaws.services.ec2.model.Image>)
	SetImages(a []*ServicesEc2ModelImage) 

	// public com.amazonaws.services.ec2.model.DescribeImagesResult withImages(com.amazonaws.services.ec2.model.Image...)
	WithImages(a ...*ServicesEc2ModelImage) *ServicesEc2ModelDescribeImagesResult

	// public com.amazonaws.services.ec2.model.DescribeImagesResult withImages(java.util.Collection<com.amazonaws.services.ec2.model.Image>)
	WithImages2(a []*ServicesEc2ModelImage) *ServicesEc2ModelDescribeImagesResult

	// public com.amazonaws.services.ec2.model.DescribeImagesResult clone()
	Clone() *ServicesEc2ModelDescribeImagesResult
}

type ServicesEc2ModelDescribeImagesResult struct {
	JavaLangObject
}

// public com.amazonaws.services.ec2.model.DescribeImagesResult()
func NewServicesEc2ModelDescribeImagesResult() (*ServicesEc2ModelDescribeImagesResult) {

	obj, err := javabind.GetEnv().NewObject("com/amazonaws/services/ec2/model/DescribeImagesResult")
	if err != nil {
		panic(err)
	}
	x := &ServicesEc2ModelDescribeImagesResult{}
	x.Callable = &javabind.Callable{obj}
	return x
}

// public java.util.List<com.amazonaws.services.ec2.model.Image> getImages()
func (jbobject *ServicesEc2ModelDescribeImagesResult) GetImages() []*ServicesEc2ModelImage {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getImages", "java/util/List")
	if err != nil {
		panic(err)
	}
	retconv := javabind.NewJavaToGoList(javabind.NewJavaToGoCallable())
	dst := new([]*ServicesEc2ModelImage)
	retconv.Dest(dst)
	if err := retconv.Convert(javabind.ObjectRef(jret)); err != nil {
		panic(err)
	}
	retconv.CleanUp()
	return *dst
}

// public void setImages(java.util.Collection<com.amazonaws.services.ec2.model.Image>)
func (jbobject *ServicesEc2ModelDescribeImagesResult) SetImages(a []*ServicesEc2ModelImage)  {
	conv_a := javabind.NewGoToJavaCollection(javabind.NewGoToJavaCallable())
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "setImages", javabind.Void, conv_a.Value().Cast("java/util/Collection"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()

}

// public com.amazonaws.services.ec2.model.DescribeImagesResult withImages(com.amazonaws.services.ec2.model.Image...)
func (jbobject *ServicesEc2ModelDescribeImagesResult) WithImages(a ...*ServicesEc2ModelImage) *ServicesEc2ModelDescribeImagesResult {
	conv_a := javabind.NewGoToJavaObjectArray(javabind.NewGoToJavaCallable(), "com/amazonaws/services/ec2/model/Image")
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "withImages", "com/amazonaws/services/ec2/model/DescribeImagesResult", conv_a.Value().Cast("com/amazonaws/services/ec2/model/Image"))
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
	unique_x := &ServicesEc2ModelDescribeImagesResult{}
	unique_x.Callable = dst
	return unique_x
}

// public com.amazonaws.services.ec2.model.DescribeImagesResult withImages(java.util.Collection<com.amazonaws.services.ec2.model.Image>)
func (jbobject *ServicesEc2ModelDescribeImagesResult) WithImages2(a []*ServicesEc2ModelImage) *ServicesEc2ModelDescribeImagesResult {
	conv_a := javabind.NewGoToJavaCollection(javabind.NewGoToJavaCallable())
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "withImages", "com/amazonaws/services/ec2/model/DescribeImagesResult", conv_a.Value().Cast("java/util/Collection"))
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
	unique_x := &ServicesEc2ModelDescribeImagesResult{}
	unique_x.Callable = dst
	return unique_x
}

// public java.lang.String toString()
func (jbobject *ServicesEc2ModelDescribeImagesResult) ToString() string {
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
func (jbobject *ServicesEc2ModelDescribeImagesResult) Equals(a interface{}) bool {
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
func (jbobject *ServicesEc2ModelDescribeImagesResult) HashCode() int {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "hashCode", javabind.Int)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

// public com.amazonaws.services.ec2.model.DescribeImagesResult clone()
func (jbobject *ServicesEc2ModelDescribeImagesResult) Clone() *ServicesEc2ModelDescribeImagesResult {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "clone", "com/amazonaws/services/ec2/model/DescribeImagesResult")
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
	unique_x := &ServicesEc2ModelDescribeImagesResult{}
	unique_x.Callable = dst
	return unique_x
}

// public java.lang.Object clone() throws java.lang.CloneNotSupportedException
func (jbobject *ServicesEc2ModelDescribeImagesResult) Clone2() (*JavaLangObject, error) {
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


