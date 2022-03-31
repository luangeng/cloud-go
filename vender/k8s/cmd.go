package k8s

func Cmd(namespace string, podID string) {
	// req := GetClient().RESTClient.Get().
	// 	Namespace(namespace).
	// 	Name(podID).
	// 	Resource("pods").
	// 	SubResource("log").
	// 	Param("follow", strconv.FormatBool(logOptions.Follow)).
	// 	Param("container", logOptions.Container).
	// 	Param("previous", strconv.FormatBool(logOptions.Previous)).
	// 	Param("timestamps", strconv.FormatBool(logOptions.Timestamps))

	// if logOptions.SinceSeconds != nil {
	// 	req.Param("sinceSeconds", strconv.FormatInt(*logOptions.SinceSeconds, 10))
	// }
	// if logOptions.SinceTime != nil {
	// 	req.Param("sinceTime", logOptions.SinceTime.Format(time.RFC3339))
	// }
	// if logOptions.LimitBytes != nil {
	// 	req.Param("limitBytes", strconv.FormatInt(*logOptions.LimitBytes, 10))
	// }
	// if logOptions.TailLines != nil {
	// 	req.Param("tailLines", strconv.FormatInt(*logOptions.TailLines, 10))
	// }
	// readCloser, err := req.Stream()
	// if err != nil {
	// 	return err
	// }

	// defer readCloser.Close()
	// _, err = io.Copy(out, readCloser)
	// return err
}
