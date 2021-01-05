package commands

//
//func addSonataInvCmd(parent *cobra.Command) {
//	parent.AddCommand(sonataInvCmd)
//
//	addFlagsForFindParams(sonataInvFindCmd)
//	sonataInvCmd.AddCommand(sonataInvFindCmd)
//
//	addFlagsForGetParams(sonataInvGetCmd)
//	sonataInvCmd.AddCommand(sonataInvGetCmd)
//}
//
//var sonataInvCmd = &cobra.Command{
//	Use:   "inventory",
//	Short: "product inventory",
//	Long:  `product inventory`,
//}
//
//var sonataInvFindCmd = &cobra.Command{
//	Use:   "find",
//	Short: "retrieve product inventory list",
//	Long:  `retrieve product inventory list`,
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
//		err = o.ExecInventoryFind(params)
//		if err != nil {
//			cmd.PrintErrln(err)
//			return
//		}
//	},
//}
//
//var sonataInvGetCmd = &cobra.Command{
//	Use:   "get",
//	Short: "retrieve product inventory item",
//	Long:  `retrieve product inventory item`,
//	Run: func(cmd *cobra.Command, args []string) {
//		var err error
//
//		params := &orchestra.GetParams{}
//		err = fillGetParamsByCmdFlags(params, cmd)
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
//		err = o.ExecInventoryGet(params)
//		if err != nil {
//			cmd.PrintErrln(err)
//			return
//		}
//	},
//}
