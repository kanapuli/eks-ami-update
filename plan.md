Check if asg_scaled count is equal to what is specified in config
Check if asg is healthy 
Check if all kubernetes nodes are healthy and its node count 
If all three is true:
	Bring up one node and move the pods to the new node and drain the old node 

If kubernetes nodes aren't enough, wait for the timeout and check again else scale asg
If asg is not healthy, scale up to bring it to a healthy state 
