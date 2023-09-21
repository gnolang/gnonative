import { GoBridge } from '@gno/native_modules';

export const initBridge = async (): Promise<boolean> => {
	try {
		console.log('bridge methods: ', Object.keys(GoBridge))

		await GoBridge.initBridge()

		return true
	} catch (err: any) {
		if (err?.message?.indexOf('already instantiated') !== -1) {
			console.log('bridge already started: ', err)
			return true
		} else {
			console.error('unable to init bridge: ', err)
		}

		return false
	}
}
