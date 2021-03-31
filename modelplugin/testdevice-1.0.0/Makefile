BASEPATH=../../../onos-helm-charts/config-models/testdevice-1.x/files
YANGPATH=${BASEPATH}/yang

models:
	go generate
	pyang -f tree \
	-p ${YANGPATH} \
	${YANGPATH}/test1@2018-02-20.yang \
	${YANGPATH}/test1-extra@2021-04-01.yang \
	-o ${BASEPATH}/testdevice-1.0.0.tree
	cat ${BASEPATH}/testdevice-1.0.0.tree