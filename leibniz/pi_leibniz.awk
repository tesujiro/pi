BEGIN{
    for(n=0;;n++){
	qpi=qpi+(-1)^n/(2*n+1)
	printf("%d:\t%f\n",n,4*qpi)
    }
}
