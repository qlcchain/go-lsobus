package commands

//
//func addSonataSiteCmd(parent *cobra.Command) {
//	parent.AddCommand(sonataSiteCmd)
//
//	addFlagsForFindParams(sonataSiteFindCmd)
//	sonataSiteCmd.AddCommand(sonataSiteFindCmd)
//}
//
//var sonataSiteCmd = &cobra.Command{
//	Use:   "site",
//	Short: "geographic site",
//	Long:  `geographic site`,
//}
//
//var sonataSiteFindCmd = &cobra.Command{
//	Use:   "find",
//	Short: "retrieve geographic site list",
//	Long:  `retrieve geographic site list`,
//	Run: func(cmd *cobra.Command, args []string) {
//		var err error
//		params := &orchestra.FindParams{}
//		err = fillFindParamsByCmdFlags(params, cmd)
//		if err != nil {
//			cmd.PrintErrln(err)
//			return
//		}
//
//		o, err := getOrchestraInstance(cmd)
//		if err != nil {
//			cmd.PrintErrln(err)
//			return
//		}
//
//		err = o.ExecSiteFind(params)
//		if err != nil {
//			cmd.PrintErrln(err)
//			return
//		}
//	},
//}
