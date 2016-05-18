package ec2

import "github.com/timob/javabind"

type ServicesEc2ModelPurchaseScheduledInstancesResultInterface interface {
	JavaLangObjectInterface

	// public java.util.List<com.amazonaws.services.ec2.model.ScheduledInstance> getScheduledInstanceSet()
	GetScheduledInstanceSet() []*ServicesEc2ModelScheduledInstance

	// public void setScheduledInstanceSet(java.util.Collection<com.amazonaws.services.ec2.model.ScheduledInstance>)
	SetScheduledInstanceSet(a []*ServicesEc2ModelScheduledInstance) 

	// public com.amazonaws.services.ec2.model.PurchaseScheduledInstancesResult withScheduledInstanceSet(com.amazonaws.services.ec2.model.ScheduledInstance...)
	WithScheduledInstanceSet(a ...*ServicesEc2ModelScheduledInstance) *ServicesEc2ModelPurchaseScheduledInstancesResult

	// public com.amazonaws.services.ec2.model.PurchaseScheduledInstancesResult withScheduledInstanceSet(java.util.Collection<com.amazonaws.services.ec2.model.ScheduledInstance>)
	WithScheduledInstanceSet2(a []*ServicesEc2ModelScheduledInstance) *ServicesEc2ModelPurchaseScheduledInstancesResult

	// public com.amazonaws.services.ec2.model.PurchaseScheduledInstancesResult clone()
	Clone() *ServicesEc2ModelPurchaseScheduledInstancesResult
}

type ServicesEc2ModelPurchaseScheduledInstancesResult struct {
	JavaLangObject
}

// public com.amazonaws.services.ec2.model.PurchaseScheduledInstancesResult()
func NewServicesEc2ModelPurchaseScheduledInstancesResult() (*ServicesEc2ModelPurchaseScheduledInstancesResult) {

	obj, err := javabind.GetEnv().NewObject("com/amazonaws/services/ec2/model/PurchaseScheduledInstancesResult")
	if err != nil {
		panic(err)
	}
	x := &ServicesEc2ModelPurchaseScheduledInstancesResult{}
	x.Callable = &javabind.Callable{obj}
	return x
}

// public java.util.List<com.amazonaws.services.ec2.model.ScheduledInstance> getScheduledInstanceSet()
func (jbobject *ServicesEc2ModelPurchaseScheduledInstancesResult) GetScheduledInstanceSet() []*ServicesEc2ModelScheduledInstance {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getScheduledInstanceSet", "java/util/List")
	if err != nil {
		panic(err)
	}
	retconv := javabind.NewJavaToGoList(javabind.NewJavaToGoCallable())
	dst := new([]*ServicesEc2ModelScheduledInstance)
	retconv.Dest(dst)
	if err := retconv.Convert(javabind.ObjectRef(jret)); err != nil {
		panic(err)
	}
	retconv.CleanUp()
	return *dst
}

// public void setScheduledInstanceSet(java.util.Collection<com.amazonaws.services.ec2.model.ScheduledInstance>)
func (jbobject *ServicesEc2ModelPurchaseScheduledInstancesResult) SetScheduledInstanceSet(a []*ServicesEc2ModelScheduledInstance)  {
	conv_a := javabind.NewGoToJavaCollection(javabind.NewGoToJavaCallable())
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "setScheduledInstanceSet", javabind.Void, conv_a.Value().Cast("java/util/Collection"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()

}

// public com.amazonaws.services.ec2.model.PurchaseScheduledInstancesResult withScheduledInstanceSet(com.amazonaws.services.ec2.model.ScheduledInstance...)
func (jbobject *ServicesEc2ModelPurchaseScheduledInstancesResult) WithScheduledInstanceSet(a ...*ServicesEc2ModelScheduledInstance) *ServicesEc2ModelPurchaseScheduledInstancesResult {
	conv_a := javabind.NewGoToJavaObjectArray(javabind.NewGoToJavaCallable(), "com/amazonaws/services/ec2/model/ScheduledInstance")
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "withScheduledInstanceSet", "com/amazonaws/services/ec2/model/PurchaseScheduledInstancesResult", conv_a.Value().Cast("com/amazonaws/services/ec2/model/ScheduledInstance"))
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
	unique_x := &ServicesEc2ModelPurchaseScheduledInstancesResult{}
	unique_x.Callable = dst
	return unique_x
}

// public com.amazonaws.services.ec2.model.PurchaseScheduledInstancesResult withScheduledInstanceSet(java.util.Collection<com.amazonaws.services.ec2.model.ScheduledInstance>)
func (jbobject *ServicesEc2ModelPurchaseScheduledInstancesResult) WithScheduledInstanceSet2(a []*ServicesEc2ModelScheduledInstance) *ServicesEc2ModelPurchaseScheduledInstancesResult {
	conv_a := javabind.NewGoToJavaCollection(javabind.NewGoToJavaCallable())
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "withScheduledInstanceSet", "com/amazonaws/services/ec2/model/PurchaseScheduledInstancesResult", conv_a.Value().Cast("java/util/Collection"))
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
	unique_x := &ServicesEc2ModelPurchaseScheduledInstancesResult{}
	unique_x.Callable = dst
	return unique_x
}

// public java.lang.String toString()
func (jbobject *ServicesEc2ModelPurchaseScheduledInstancesResult) ToString() string {
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
func (jbobject *ServicesEc2ModelPurchaseScheduledInstancesResult) Equals(a interface{}) bool {
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
func (jbobject *ServicesEc2ModelPurchaseScheduledInstancesResult) HashCode() int {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "hashCode", javabind.Int)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

// public com.amazonaws.services.ec2.model.PurchaseScheduledInstancesResult clone()
func (jbobject *ServicesEc2ModelPurchaseScheduledInstancesResult) Clone() *ServicesEc2ModelPurchaseScheduledInstancesResult {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "clone", "com/amazonaws/services/ec2/model/PurchaseScheduledInstancesResult")
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
	unique_x := &ServicesEc2ModelPurchaseScheduledInstancesResult{}
	unique_x.Callable = dst
	return unique_x
}

// public java.lang.Object clone() throws java.lang.CloneNotSupportedException
func (jbobject *ServicesEc2ModelPurchaseScheduledInstancesResult) Clone2() (*JavaLangObject, error) {
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


