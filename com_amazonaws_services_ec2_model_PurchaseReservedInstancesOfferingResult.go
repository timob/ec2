package ec2

import "github.com/timob/javabind"

type ServicesEc2ModelPurchaseReservedInstancesOfferingResultInterface interface {
	JavaLangObjectInterface

	// public void setReservedInstancesId(java.lang.String)
	SetReservedInstancesId(a string) 

	// public java.lang.String getReservedInstancesId()
	GetReservedInstancesId() string

	// public com.amazonaws.services.ec2.model.PurchaseReservedInstancesOfferingResult withReservedInstancesId(java.lang.String)
	WithReservedInstancesId(a string) *ServicesEc2ModelPurchaseReservedInstancesOfferingResult

	// public com.amazonaws.services.ec2.model.PurchaseReservedInstancesOfferingResult clone()
	Clone() *ServicesEc2ModelPurchaseReservedInstancesOfferingResult
}

type ServicesEc2ModelPurchaseReservedInstancesOfferingResult struct {
	JavaLangObject
}

// public com.amazonaws.services.ec2.model.PurchaseReservedInstancesOfferingResult()
func NewServicesEc2ModelPurchaseReservedInstancesOfferingResult() (*ServicesEc2ModelPurchaseReservedInstancesOfferingResult) {

	obj, err := javabind.GetEnv().NewObject("com/amazonaws/services/ec2/model/PurchaseReservedInstancesOfferingResult")
	if err != nil {
		panic(err)
	}
	x := &ServicesEc2ModelPurchaseReservedInstancesOfferingResult{}
	x.Callable = &javabind.Callable{obj}
	return x
}

// public void setReservedInstancesId(java.lang.String)
func (jbobject *ServicesEc2ModelPurchaseReservedInstancesOfferingResult) SetReservedInstancesId(a string)  {
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
func (jbobject *ServicesEc2ModelPurchaseReservedInstancesOfferingResult) GetReservedInstancesId() string {
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

// public com.amazonaws.services.ec2.model.PurchaseReservedInstancesOfferingResult withReservedInstancesId(java.lang.String)
func (jbobject *ServicesEc2ModelPurchaseReservedInstancesOfferingResult) WithReservedInstancesId(a string) *ServicesEc2ModelPurchaseReservedInstancesOfferingResult {
	conv_a := javabind.NewGoToJavaString()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "withReservedInstancesId", "com/amazonaws/services/ec2/model/PurchaseReservedInstancesOfferingResult", conv_a.Value().Cast("java/lang/String"))
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
	unique_x := &ServicesEc2ModelPurchaseReservedInstancesOfferingResult{}
	unique_x.Callable = dst
	return unique_x
}

// public java.lang.String toString()
func (jbobject *ServicesEc2ModelPurchaseReservedInstancesOfferingResult) ToString() string {
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
func (jbobject *ServicesEc2ModelPurchaseReservedInstancesOfferingResult) Equals(a interface{}) bool {
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
func (jbobject *ServicesEc2ModelPurchaseReservedInstancesOfferingResult) HashCode() int {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "hashCode", javabind.Int)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

// public com.amazonaws.services.ec2.model.PurchaseReservedInstancesOfferingResult clone()
func (jbobject *ServicesEc2ModelPurchaseReservedInstancesOfferingResult) Clone() *ServicesEc2ModelPurchaseReservedInstancesOfferingResult {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "clone", "com/amazonaws/services/ec2/model/PurchaseReservedInstancesOfferingResult")
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
	unique_x := &ServicesEc2ModelPurchaseReservedInstancesOfferingResult{}
	unique_x.Callable = dst
	return unique_x
}

// public java.lang.Object clone() throws java.lang.CloneNotSupportedException
func (jbobject *ServicesEc2ModelPurchaseReservedInstancesOfferingResult) Clone2() (*JavaLangObject, error) {
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


