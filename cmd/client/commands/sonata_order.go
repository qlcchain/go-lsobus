package commands

//
//func addSonataOrderCmd(parent *cobra.Command) {
//	parent.AddCommand(sonataOrderCmd)
//
//	addFlagsForOrderParams(sonataOrderCreateCmd)
//	sonataOrderCmd.AddCommand(sonataOrderCreateCmd)
//
//	addFlagsForFindParams(sonataOrderFindCmd)
//	sonataOrderCmd.AddCommand(sonataOrderFindCmd)
//
//	addFlagsForGetParams(sonataOrderGetCmd)
//	sonataOrderCmd.AddCommand(sonataOrderGetCmd)
//}
//
//var sonataOrderCmd = &cobra.Command{
//	Use:   "order",
//	Short: "product ordering",
//	Long:  `product ordering`,
//}
//
//var sonataOrderCreateCmd = &cobra.Command{
//	Use:   "create",
//	Short: "create product ordering",
//	Long:  `create product ordering`,
//	Run: func(cmd *cobra.Command, args []string) {
//		var err error
//
//		params := &orchestra.OrderParams{}
//		err = fillOrderParamsByCmdFlags(params, cmd)
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
//		err = o.ExecOrderCreate(params)
//		if err != nil {
//			cmd.PrintErrln(err)
//			return
//		}
//	},
//}
//
//var sonataOrderFindCmd = &cobra.Command{
//	Use:   "find",
//	Short: "retrieve product ordering list",
//	Long:  `retrieve product ordering list`,
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
//		err = o.ExecOrderFind(params)
//		if err != nil {
//			cmd.PrintErrln(err)
//			return
//		}
//	},
//}
//
//var sonataOrderGetCmd = &cobra.Command{
//	Use:   "get",
//	Short: "retrieve product ordering item",
//	Long:  `retrieve product ordering item`,
//	Run: func(cmd *cobra.Command, args []string) {
//		var err error
//		op := &orchestra.GetParams{}
//		err = fillGetParamsByCmdFlags(op, cmd)
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
//		o.ExecOrderGet(op)
//	},
//}
