package main

import (
	"fmt"
)

// method

type Students struct{

    Name string
    Age int32
    Grades []int32
    
}

func addStd(name *string,age *int32, grds *[]int32) (string,int32,[]int32) {
        
        var numGrades int

        fmt.Println("Enter Student Name : ")
        fmt.Scanln(name)

        fmt.Println("Enter Student Age : ")
        fmt.Scanln(age)
        
        fmt.Println("No of Grades : ")
        fmt.Scanln(&numGrades)

        fmt.Println("Enter markes : ")

        for i:=0;i< numGrades;i++ {

            var mark int32

            fmt.Scanln(&mark)
            *grds = append(*grds , mark)

        }

        return *name,*age,*grds
}

func (s_val *Students) toChangeStudent() {
    var cp_name string
    
    fmt.Println("Enter Name : ")
    fmt.Scanln(&cp_name)

    s_val.Name = cp_name
    
}

func main()  {

    var opt int
    var name string
    var age int32
    var grds []int32
    var edit string

    fmt.Println("\t WELCOME \n\t Do you want to enter student Details [Yes: 1, No: Any other key]")
    fmt.Scanln(&opt)

    if opt ==1{

        studentName, studentAge, studentGrades :=addStd(&name,&age,&grds)

        student := Students{
            Name: studentName,
            Age: studentAge,
            Grades: studentGrades,
        }

        fmt.Println("Before : ",student)

        fmt.Println("Enter 'edit' to edit student info or any key to No :")
        fmt.Scanln(&edit)

        if edit == "edit" {
            student.toChangeStudent()
            fmt.Println("After : ",student)
        }else {
            fmt.Println("No edits made.")
        }
        // fmt.Println(student)

    }else {
        fmt.Println("Exiting. ")
    }
}
