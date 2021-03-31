BASEPATH=../../../onos-helm-charts/config-models/testdevice-2.x/files
YANGPATH=${BASEPATH}/yang

models:
	go generate
	pyang -f tree \
	-p ${YANGPATH} \
	${YANGPATH}/test1@2019-06-10.yang \
	${YANGPATH}/test1-augmented@2020-02-29.yang \
	-o ${BASEPATH}/testdevice-2.0.0.tree
	cat ${BASEPATH}/testdevice-2.0.0.tree