

def class_html():
    course_name = input("Course Name: ")
    labs = input("Number of Labs: ")
    low_student_number = input("Lowest Student ID Number: ")
    high_student_number = input("Highest Student ID Number: ")
    print("Your Class Link is http://alta3.com/graphs?={}&={}&={}&={}".format(
        course_name, labs, low_student_number, high_student_number)

class_html()
