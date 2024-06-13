package main_test

import (
	main "a21hc3NpZ25tZW50"
	"a21hc3NpZ25tZW50/model"
	"time"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("Main", func() {
	var sm *main.InMemoryStudentManager

	BeforeEach(func() {
		sm = main.NewInMemoryStudentManager(3)
	})

	Describe("GetStudents", func() {
		It("should return all students", func() {
			students := sm.GetStudents()
			Expect(students).To(HaveLen(3))
			Expect(students[0].ID).To(Equal("A12345"))
			Expect(students[0].Name).To(Equal("Aditira"))
			Expect(students[0].StudyProgram).To(Equal("TI"))
			Expect(students[1].ID).To(Equal("B21313"))
			Expect(students[1].Name).To(Equal("Dito"))
			Expect(students[1].StudyProgram).To(Equal("TK"))
			Expect(students[2].ID).To(Equal("A34555"))
			Expect(students[2].Name).To(Equal("Afis"))
			Expect(students[2].StudyProgram).To(Equal("MI"))
		})
	})

	Describe("Login", func() {
		When("the ID and name match a student record", func() {
			It("should return success message", func() {
				msg, err := sm.Login("A12345", "Aditira")
				Expect(msg).To(Equal("Login berhasil: Selamat datang Aditira! Kamu terdaftar di program studi: Teknik Informatika"))
				Expect(err).ToNot(HaveOccurred())
			})
		})

		When("the ID is empty", func() {
			It("should return an error", func() {
				msg, err := sm.Login("", "Aditira")
				Expect(msg).To(Equal(""))
				Expect(err).To(HaveOccurred())
			})
		})

		When("the name is empty", func() {
			It("should return an error", func() {
				msg, err := sm.Login("A12345", "")
				Expect(msg).To(Equal(""))
				Expect(err).To(HaveOccurred())
			})
		})

		When("the ID and name do not match any student record", func() {
			It("should return an error", func() {
				msg, err := sm.Login("invalid_id", "invalid_name")
				Expect(msg).To(Equal(""))
				Expect(err).To(HaveOccurred())
			})
		})

		When("an invalid ID and name are used more than 3 times", func() {
			It("should return an error after 3 attempts", func() {
				var msg string
				var err error

				for i := 0; i < 3; i++ {
					msg, err = sm.Login("invalid_id", "invalid_name")
				}

				Expect(msg).To(Equal(""))
				Expect(err).To(HaveOccurred())
				Expect(err.Error()).To(Equal("Login gagal: data mahasiswa tidak ditemukan"))

				// 4th attempt
				msg, err = sm.Login("invalid_id", "invalid_name")
				Expect(msg).To(Equal(""))
				Expect(err).To(HaveOccurred())
				Expect(err.Error()).To(Equal("Login gagal: Batas maksimum login terlampaui"))
			})

			It("should reset the count after a successful login", func() {
				// Fail login 2 times
				for i := 0; i < 2; i++ {
					msg, err := sm.Login("A12345", "invalid_name")
					Expect(msg).To(Equal(""))
					Expect(err).To(HaveOccurred())
					Expect(err.Error()).To(Equal("Login gagal: data mahasiswa tidak ditemukan"))
				}

				// Login successfully with a valid ID and name
				msg, err := sm.Login("A12345", "Aditira")
				Expect(msg).To(Equal("Login berhasil: Selamat datang Aditira! Kamu terdaftar di program studi: Teknik Informatika"))
				Expect(err).ToNot(HaveOccurred())

				// Attempt to fail login 3 times again with the same ID, should be blocked on the 4th attempt
				for i := 0; i < 3; i++ {
					msg, err = sm.Login("A12345", "invalid_name")
					Expect(msg).To(Equal(""))
					Expect(err).To(HaveOccurred())
					Expect(err.Error()).To(Equal("Login gagal: data mahasiswa tidak ditemukan"))
				}

				msg, err = sm.Login("A12345", "invalid_name")
				Expect(msg).To(Equal(""))
				Expect(err).To(HaveOccurred())
				Expect(err.Error()).To(Equal("Login gagal: Batas maksimum login terlampaui"))
			})
		})
	})

	Describe("Register", func() {
		When("all required fields are provided and ID is unique", func() {
			It("should add the student to the list", func() {
				msg, err := sm.Register("C12345", "Citra", "SI")
				Expect(msg).To(Equal("Registrasi berhasil: Citra (SI)"))
				Expect(err).ToNot(HaveOccurred())

				students := sm.GetStudents()
				Expect(students).To(HaveLen(4))
				Expect(students[3].ID).To(Equal("C12345"))
				Expect(students[3].Name).To(Equal("Citra"))
				Expect(students[3].StudyProgram).To(Equal("SI"))
			})
		})

		When("ID is already used", func() {
			It("should return an error", func() {
				msg, err := sm.Register("A12345", "Aditira", "TI")
				Expect(msg).To(Equal(""))
				Expect(err).To(HaveOccurred())
				Expect(err.Error()).To(Equal("Registrasi gagal: id sudah digunakan"))
			})
		})

		When("ID, Name or StudyProgram is empty", func() {
			It("should return an error", func() {
				msg, err := sm.Register("", "", "TK")
				Expect(msg).To(Equal(""))
				Expect(err).To(HaveOccurred())
				Expect(err.Error()).To(Equal("ID, Name or StudyProgram is undefined!"))
			})
		})

		When("StudyProgram is invalid", func() {
			It("should return an error", func() {
				msg, err := sm.Register("C12345", "Citra", "ABC")
				Expect(msg).To(Equal(""))
				Expect(err).To(HaveOccurred())
				Expect(err.Error()).To(Equal("Study program ABC is not found"))
			})
		})
	})

	Describe("GetStudyProgram", func() {
		When("given an existing study program code", func() {
			It("should return the correct study program name", func() {
				name, err := sm.GetStudyProgram("TI")
				Expect(err).NotTo(HaveOccurred())
				Expect(name).To(Equal("Teknik Informatika"))
			})
		})

		When("given an undefined study program code", func() {
			It("should return an error", func() {
				name, err := sm.GetStudyProgram("")
				Expect(err).To(HaveOccurred())
				Expect(name).To(BeEmpty())
			})
		})

		When("given a non-existing study program code", func() {
			It("should return an error", func() {
				name, err := sm.GetStudyProgram("unknown")
				Expect(err).To(HaveOccurred())
				Expect(name).To(BeEmpty())
			})
		})
	})

	Describe("ModifyStudent", func() {
		When("given an existing student name", func() {
			It("should modify the student's study program", func() {
				modifier := sm.ChangeStudyProgram("SI")

				msg, err := sm.ModifyStudent("Afis", modifier)
				Expect(err).NotTo(HaveOccurred())
				Expect(msg).To(Equal("Program studi mahasiswa berhasil diubah."))

				name, err := sm.GetStudyProgram("SI")
				Expect(err).NotTo(HaveOccurred())
				Expect(name).To(Equal("Sistem Informasi"))
			})
		})

		When("given a non-existing student name", func() {
			It("should return an error", func() {
				modifier := sm.ChangeStudyProgram("SI")

				msg, err := sm.ModifyStudent("unknown", modifier)
				Expect(err).To(HaveOccurred())
				Expect(msg).To(BeEmpty())
			})
		})
	})

	Describe("ImportStudents", func() {
		When("given several large valid CSV files", func() {
			It("should import the students quickly, less than 300ms, but more than 50ms", func() {
				// Paths to your existing CSV files
				filepaths := []string{"students1.csv", "students2.csv", "students3.csv"}

				// Measure the time it takes to import the students
				start := time.Now()
				err := sm.ImportStudents(filepaths)
				elapsed := time.Since(start)

				// Check that the students were imported without errors
				Expect(err).NotTo(HaveOccurred())

				// Check that the import was fast
				Expect(elapsed).To(BeNumerically("<", 300*time.Millisecond))

				// Check that the import was not too fast
				Expect(elapsed).To(BeNumerically(">", 50*time.Millisecond))
			})

			It("should correctly import the students", func() {
				// Paths to your existing CSV files
				filepaths := []string{"students1.csv", "students2.csv", "students3.csv"}

				// Import the students from the CSV files
				err := sm.ImportStudents(filepaths)

				// Check that the students were imported without errors
				Expect(err).NotTo(HaveOccurred())

				// Check that the students were imported
				students := sm.GetStudents()
				Expect(len(students)).To(BeNumerically(">=", 3000)) // replace with the expected number of students

				// Check specific students
				Expect(containsStudent(students, "H73886")).To(BeTrue())
				Expect(containsStudent(students, "Y00313")).To(BeTrue())
			})
		})
	})

	Describe("SubmitAssignment", func() {
		When("student submit assignments", func() {
			It("should submit the assignment quickly, less than 100ms", func() {

				start := time.Now()
				sm.SubmitAssignments(10)
				elapsed := time.Since(start)

				Expect(elapsed).To(BeNumerically(">", 110*time.Millisecond))

				Expect(elapsed).To(BeNumerically("<", 200*time.Millisecond))

			})
		})
	})
})

// Helper function to check if a slice of students contains a student with a specific ID
func containsStudent(students []model.Student, id string) bool {
	for _, student := range students {
		if student.ID == id {
			return true
		}
	}
	return false
}
