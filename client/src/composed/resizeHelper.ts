import { reactive } from "vue"

const resizeHelper = () => {
    const data = reactive<{
        info: {
            initialWidth: number,
            initialHeight: number,
            width: number,
            height: number,
            el : HTMLDivElement | null
        },
        
        coordinateInfo : {
            initialCoordinateX: number,
            initialCoordinateY: number,
            coordinateX: number,
            coordinateY: number,
        },
        hookAfterCoordinateChanges() : void
    }>({
        info: {
            initialWidth: 0,
            initialHeight: 0,
            width: 0,
            height: 0,
            el: null
        },
        coordinateInfo: {
            initialCoordinateX: 0,
            initialCoordinateY: 0,
            coordinateX: 0,
            coordinateY: 0
        },
        hookAfterCoordinateChanges: () => {}
    })

    const dragStart = (event : DragEvent) => {
        event.dataTransfer!.setDragImage(document.createElement('div'), 0, 0);
        data.info.initialWidth = data.info.width
        data.info.initialHeight = data.info.height
        data.coordinateInfo.initialCoordinateX = event.clientX;
        data.coordinateInfo.initialCoordinateY = event.clientY;
    }

    const dragOver = (event: DragEvent) => {
        event.preventDefault()
    }
    const drag = (event : DragEvent) => {
        event.preventDefault()
        if (event.screenX == 0 && event.screenY == 0) {
            return;
        }

        data.coordinateInfo.coordinateX = event.clientX;
        data.coordinateInfo.coordinateY = event.clientY
        const diffX = data.coordinateInfo.coordinateX - data.coordinateInfo.initialCoordinateX;
        const diffY = data.coordinateInfo.coordinateY - data.coordinateInfo.initialCoordinateY;

        data.info.width = data.info.initialWidth + diffX;
        data.info.height = data.info.initialHeight + diffY;
        data.hookAfterCoordinateChanges()
    }

    const set = (el : HTMLDivElement ,option : {
        initialWidth: number,
        initialHeight: number,
        hookAfterCoordinateChanges?: () => void
    }) => {
        el.draggable = true
        data.info.el = el
        data.info.initialWidth = option.initialWidth
        data.info.initialHeight = option.initialHeight
        data.info.width = option.initialHeight
        data.info.height = option.initialHeight

        data.info.el.addEventListener('dragstart', dragStart);
        data.info.el.addEventListener('dragover', dragOver);
        data.info.el.addEventListener('drag', drag);
        if (option.hookAfterCoordinateChanges) {
            data.hookAfterCoordinateChanges = option.hookAfterCoordinateChanges
        }
        
    }
    const deleteListeners = () => {
        if (data.info.el) {
            data.info.el.removeEventListener('dragstart', dragStart);
        data.info.el.removeEventListener('dragover', dragOver);
        data.info.el.removeEventListener('drag', drag);
        }
    }

    return {
        data,
        set,
        deleteListeners
    }
}

export default resizeHelper