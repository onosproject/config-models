BASEPATH=../../../onos-helm-charts/config-models/devicesim/files
YANGPATH=${BASEPATH}/yang

models:
	go generate
	pyang -f tree \
	-p ${YANGPATH} \
	${YANGPATH}/openconfig-interfaces@2017-07-14.yang \
	${YANGPATH}/openconfig-openflow@2017-06-01.yang \
	${YANGPATH}/openconfig-platform@2016-12-22.yang \
	${YANGPATH}/openconfig-system@2017-07-06.yang \
	-o ${BASEPATH}/devicesim-1.0.0.tree
	cat ${BASEPATH}/devicesim-1.0.0.tree