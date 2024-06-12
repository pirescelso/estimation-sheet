package domain_test

// func ExampleItem() {
// 	originalYear := 2020
// 	originalMonth := time.January
// 	plannedStart := time.Date(originalYear, originalMonth, 1, 0, 0, 0, 0, time.UTC)
//
// 	currency := cost.NewCurrency("USD", "$", "United States of America")
// 	installments := []cost.Installment{
// 		cost.NewInstallment(originalYear, originalMonth+1, 100),
// 		cost.NewInstallment(originalYear, originalMonth+7, 200),
// 	}
// 	item := cost.NewItem(
// 		"Mão de obra do PO",
// 		"estimativa do Ferraz",
// 		cost.NewOneTimeCost(
// 			300,
// 			currency,
// 			installments,
// 		),
// 	)
//
// 	project := cost.NewProject(plannedStart, item)
// 	output(project, "Original")
// 	fmt.Println("--------------------")
// 	if err := project.AddMonths(12); err != nil {
// 		fmt.Printf("STOP: %s\n", err)
// 		return
// 	}
// 	output(project, "Changed")
//
// 	// Output:
// 	// STOP
// }
//
// func output(project *cost.Project, title string) {
// 	fmt.Printf("%s:\n PlannedStart: %s\n", title, project.PlannedStart.Format("2006-01-02"))
// 	for _, v := range project.Items {
// 		for j, v := range v.Installments {
// 			fmt.Printf("   Installment %d: %s %f\n", j, v.Date.Format("2006-01-02"), v.Amount)
// 		}
// 	}
// }

// func TestProject(t *testing.T) {
// 	t.Run("should create a project with valid values", func(t *testing.T) {
// 		plannedStart := time.Date(2020, time.January, 1, 0, 0, 0, 0, time.UTC)
// 		project := domain.NewProject(plannedStart)

// 		assert.Equal(t, plannedStart, project.PlannedStart)
// 	})

// 	t.Run("should add items", func(t *testing.T) {
// 		project := arrangeProject()
// 		otc := arrangeItemWithInstallment(OTC)
// 		rc := arrangeItemWithInstallment(RC)
// 		inv := arrangeItemWithInstallment(INV)
// 		project.AddItems(otc, rc, inv)

// 		assert.Equal(t, 3, len(project.Items))
// 	})

// 	t.Run("should add months", func(t *testing.T) {
// 		project := arrangeProjectWithItems()
// 		projectClone := arrangeProjectWithItems()

// 		project.AddMonths(5)

// 		assert.Equal(t, time.Duration(3648*time.Hour), project.PlannedStart.Sub(projectClone.PlannedStart))

// 		for i := 0; i < len(project.Items); i++ {
// 			new := project.Items[i].GetInstalmments()
// 			old := projectClone.Items[i].GetInstalmments()
// 			for j := 0; j < len(new); j++ {
// 				assert.NotEqual(t, old[j].Date, new[j].Date)
// 				top := time.Duration(3672 * time.Hour)
// 				assert.WithinDuration(t, new[j].Date, old[j].Date, top)

// 			}
// 		}
// 	})
// }
