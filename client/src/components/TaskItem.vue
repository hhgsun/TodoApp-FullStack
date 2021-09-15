<template>
  <div class="task-item" :class="task.Done ? 'done' : ''">
    <button class="btn-done" @click="doneItem(task)" :disabled="isWait">
      <img
        v-if="task.Done"
        src="data:image/png;base64,iVBORw0KGgoAAAANSUhEUgAAADIAAAAyCAYAAAAeP4ixAAAABmJLR0QA/wD/AP+gvaeTAAAD3UlEQVRoge3ZSYhcRRzH8Y+ZQZsxLoijOSiMAcElHkxQIxokIgFBIRi9COpNJAsB8aIRFDyMB/GgqDcvMYIiIrjeRAVxjczEGBUxOSg5RIPLJJkZ4rSH6kdXv17mva6aPuUHBfXoqv///+16tf0fZ3RGK6KzMtlpYCNuwTVYi9WYwN/4B79iBl/iKyxl8p1Fm7EXc2jWKEfxMq4dfcid2iL8s3WC71WW8B7WpQZU99WaxEu4r8dvc/gU+/EzjmMB5+Ei4ZXbgJtxdqnvabyAJ3GqZky1tQm/6/5HP8DWHsH104V4EN/qHqH9uGJA3/FhAo+1DfMlp5/g+kS7dwqjF9v9A+t7tF2DV1KcbROGvnB0CtvlXfGe0wlzHDdEbdbgED4e1skmnSNxDDcOa2wZ3V/ydRSXaUM0DQkyqXNOHMPV6fEO1N1YjHx+rQ0xNMibOl+nlRqJsh7SvQgMDbKlZGB7tjCr6VWZQD6POn8m38SuokvxvQwgm6OOS9KX2DqKJ3YyyN6o4/tZwxys5SBqgTR0HgC3Zg62n6pADAQpb/kbcW6r/q9w/FhpjWMPfmuVQZqpanSPNv0oILJpVek53vC+GWUgqSqDxCfPX0YZSKrKIBdE9T9HGUiqyiCNqD4/ykBSVQY5GdUnRhlIqsogf0X1ycy+xrAbB4QrcMpdfwGz2NWy26X4sPZsZoh3EoPvV97uBfNY1GDo21gP7V4hiKLsLJ9qb8IXrfo8LsaJDCAHtFM+7+IJnfOxriYwjbtaz7PlBuPCsluQPpDgLFY8J5JzWC2ti2wulCf7abwRPT+SyWmcKprLZDO20zMVtV7n+3d7BqexvakM9rTsxHZ76qOowaz0xFgqSEP3dlAJZIPOXNYzQziPlQLSwIc4KFyDC1UCgRejRv8JqZphNQjknAH9Coii70Ehj0wNkAmdSYATuK1W+G31A7lEWJp39oGIX/GmcA0vNr/KIIT7Sbwcn8Q99Tl6ghQQTSHJsaMGRG0Qwleo+B6/JORqG4M6VQCZwpGS3R0VIYYCIdzl45FpCheve3UfPKuC9IMp57Re6wExNAhcie9KnZv4AY8KCee6IIRb6RHddgdBJIEQhn1a/2P4j9iHp4QJ/HCrLLf89oLZNwAiGaTQVXhd515TtfTbR2KY5SCygRRai6eFkUgFIcA8b3kISiA5k9OX41bt7+zna69sd0Ttin8+VVM4nMFOLaUcUfqpY0SqLJ05tBjVV2eyGdtZTP7cW1E/4bpWfRqPy3NDLHQowVYt7VJ/datTRvZFbUzIdqwExFuqnS6yaUzYJGflyWvNCCOxCv4HZU8tbNCX4SUAAAAASUVORK5CYII="
      />
      <img
        v-else
        src="data:image/png;base64,iVBORw0KGgoAAAANSUhEUgAAADIAAAAyCAYAAAAeP4ixAAAABmJLR0QA/wD/AP+gvaeTAAAEjklEQVRoge3aXYgWVRgH8N9q67tWmra6akIJdpGZESVSZEQUmmUqhlBpFxVaSp9SBAWadFNEF5HUVVB05UVdhJJSpGWRmERZaebShWlRZEWGX9luF+dMZ2zf3Z1539k1wj8M7+75eJ7/zJxznv9zznAa/y20VGhrNK7F1bgYk9GBs6KfP/AzOrELH2IrDlbIoWHUsBibcALdJa8/sRFLMGyQuYPhWIkDOVJHsRmrsABThLfUGq/RwptagDXYguO5/geizUG7oTnC8MgIfIplGNWArXbchy9z9jpxYyVMe0EbXso5/Ey4qSrmWQvmYXe03YWXo89KMQ47opMjeARnVO1EGIJP4Fj09THG5+obeev/YBL2RsN7cVkzxgriCuyL14RYthQ/YEYjBsdJN7EdY5vnWBgTMDX+fb20MNxV1lCbNJy2Y0RFBMviIvwSeTzTiIFsYu81uG8ijzHSCvkGhpQ1MEea2IMxJ+qhhg8ijx04M5ZPLGpguPQUVlbNriBa8Frk8B3Oi+X34CsFV8yVUpwYiCW2CJ6MHA45eUR8EcuX9megJsmOOQNAsAgWCQHxLyFI5nG7JGda+zKyWJIdVSrjopiBw3of1kMkOfPvmzwJm2KjZRUTLIKJ2B/9v9JHu2zov9lbg9GCFD+qSSnQAEbgc4HgO/oeNhMEnsdxTr0GC6Kh96rl2C+GYn30vVuxh5gty7OygnyAmRl/t1REsCiex81C9jgXvxXosy3+XlWvMnsq86tgVxDLpaRsZj9t81gU+22oV/lNrJzSLLuCmC2kul24s2TfS6VY1wMHY+W5BY0NxTrcXZIE4WH9Gv093UD/iVLU74EsmSmaM98qZXMPlCDRgW9j33Uai1fDY//D9SrL3gjcL0TgohK7DR/pKQTLoibNrR4oO7QyLBHGejfW6l1qt+D12G6/Eiq2DsZGOz/Vq9yj8cm+UHg63XhVfbG5Otb/jmkN+MhjqhR3emCD5pbf6wSS3XjLyTsgmRA8gVsatJ/HwpyfHnguVq5qwsGVUlr6Ls4WgtaRWPZgE7bzWBPtPVuvcn6s3Nykk2n4Xsr1f5TmT1XIJMrcepWjJNE4uklHk6Ulthtvqy5JaxcE43GM7K3Rxuj43gocjhcU7S7VqukVAsf1fTW6Q7WJVTvOr8BOhhZBlnTjtr4a1qTk5qYKCVSFeZI06TdwPxwb73TqNh/qoVVKcwtJojZpm/TRgeNVGo8LnL4WRk4hzJa0zOUDw6sUpgtcunBD2c5rpYOXjmp5lcI4aSl/oREDNSGgdeMTp2YTe6S0kb5NE8dyY4QxmcnuwXwz+cOlzvh/U5gkpcGdBmfOTJeG0x5cUJXhDmmYHcVj+tmybBCtwuqUpQTbDMCRRg0vSvpppyDaqjwMzeJElzCxB/SoepY01LKdjOXKZ5YECbNCkh1ZnCi9xDaKGh4SpEJG4Jggr1cLic8lkeiweLXHsoV4Svh8I//BwD4hYhcOdlWiJmz1b5Dy9rKfcKwXBGBTw6jKo4ORwkc11wh5/4XCRM3izyFhs6BTyLW34v1Yfhr/O/wNK7RXqMq67vYAAAAASUVORK5CYII="
      />
    </button>
    <div v-if="isEditable" class="editable">
      <input v-model="currentText" />
      <button v-if="isEditable" class="btn-edit" @click="updateText(task)">
        <img
          src="data:image/png;base64,iVBORw0KGgoAAAANSUhEUgAAADIAAAAyCAYAAAAeP4ixAAAABmJLR0QA/wD/AP+gvaeTAAABXElEQVRoge2YS1ICMRBAn5aFx9E1eD0/HEV3bLwAehOHjZ5BXBBKoDLm051Jl/SrSg1Uejp5hGSgwXGcUmbAEtgA207tOcxDxFNHgcO2ksoMIdF8pH8/UAv2ub9QkElNdAqRG+AzvH4FriXJavtTvHP8FVqP5L5FuDKtRWL7YSy3aGWmEonliuWulrEmApUyFkXgeM+8lA5U01+Sv0QEdjJb4Lt0oJr+kvylIqMxl4IJSVF9FvUQeTt5v45GKWDhyV4cU7Mim3BdVNz7F3fhOmglTH0qjwcxLdq9cH7ZgbMgM6Ar8BEkUr+p1ER6Y+74VcVFrHHWIi2rLAO7E1FcQYH0qTVFleVBML/swFSVRcKC35UZQ02k9XOmavyz3uwmcRFruIg1XMQaLmKNHJHef32zxv83K3KVEXPRfBYK48dWpFUBToOiIl7rApxGSxXxgHYFOI2WW8RzHOeEH/58jcLJRDukAAAAAElFTkSuQmCC"
        />
      </button>
    </div>
    <div v-else class="text">
      {{ task.Text }}
    </div>
    <small v-if="!isEditable" class="date">{{
      dateFilter(task.UpdatedAt)
    }}</small>
    <div class="task-actions">
      <button
        v-if="!isEditable && !task.Done"
        class="btn-edit"
        @click="openEditable()"
      >
        <img
          src="data:image/png;base64,iVBORw0KGgoAAAANSUhEUgAAADIAAAAyCAYAAAAeP4ixAAAABmJLR0QA/wD/AP+gvaeTAAACX0lEQVRoge3YO2sUURjG8V/cVYKK6CewE1sFERVvGC8RRSy8omLhBxBsbMUbghYpLEQQ75ooFtaSSlFBe60liI0p1Jh4L84ZZlw3u7MhuzuG+cPLLO857+7zzDl7zsyhpKSk5H+ip0u/uxzHsBIL8B5PcAUfuqSpJSoYwC/8rhOfcKBr6nJSxZAg+BsuY5MwOnsxHNt+YU+XNDalikHpnT9Vp08Pzsb2j1jYMXU5yZr4Ga8T2FWnbw+exj7HOyUwD1XcF4R9QR+uSs3srFNzNLY/7pDGplRxT2piU8zP8reZHTV1a2Pbq87IbExFfRMJWTPj2J5p2x3zw+2X2ZgK7prcREKtmf6YvxVz59srszEV3NHcRELWzFecFhaEH1jaPpmNqeC2/CYSsmaSuNAOgXmYqomEPunSPBS/r+NUcDOKGNO6iXX4HOsfYPa0qstJBTekJvparF9vBpjYHOu6amKWGWLiuqmb2KI0MX1cM/XVaauw6SVLbHV6peUneZj7jTMt1m5TEBNwUmrkBw7mrMuORFenU8IjQcybeP0uvKI2ol94ICzESCSMCIJW4Zx0ZA5N0j87nQoxErBY+iLUG3NZM4dr+hfSBOwTRL2oyWfNHIm5wpqAS4KwgTptF6VmBhTwP5HlmSBuf02+F2ukC0AShTQxRzpVNggngQPCNJvw72nhoAKagBXqH28m8Q4PcQKrdek8Oc+dm5v5PIbXeInn8TrSBl1tYyOWKeiUKSkpaU6zpXJRjCIwGqNl5sfCRntIJ2MU8yYT22g5HRd27yWtuG8jb4UniZKSkpIZwB9Nu/kCteMO3AAAAABJRU5ErkJggg=="
        />
      </button>
      <button
        v-if="!isEditable"
        class="btn-delete"
        @click="removeItem(task)"
        :disabled="isWait"
      >
        <img
          src="data:image/png;base64,iVBORw0KGgoAAAANSUhEUgAAADIAAAAyCAYAAAAeP4ixAAAABmJLR0QA/wD/AP+gvaeTAAACLElEQVRoge2ZPU+UQRCAn+MIZ+TDUCCS8NEgMfIDMKEjhobCBhMs6PwFR01FD4GSPyCVsbFQ7K34AVYaPgTuYkFOAoiXo5g9LiRwt7Pv7uYC+ySbKW5nZmdnv25eSCQSD4KcZ3s9wAowCzy6o8858AVYBk49+/fGFlCzbB98OvaZkcdABagCr4H9O/qNANtAHsngmccxeGEMmemSRd+S6Tvmy3lni9+7kdntsrA1YGQBeNuib8HIOaBsYfsf8I0Me2oD+zUfuq03G2irjHwCnlv0C81/M5ZEIuGI9kIsAq9CDOQWvgNroYz/Jt5xe6AZmPZYPQCGgFXgh1LXlglgCbsXgjMfkdl6F9DHvPHxWaPUoXSyZ+SoUk/DMyOPNEqugYwo9TRECWTXyJCBDBkZJSMxltahRqkdA3HKiJY88t+gBvQG8lG/q8YD2b/ml3H0MoDtDuDS2O/RKmoJeXI9RS7pCvBXo5glkBD7xOnoBbdAQh7Bzhu9XTOiOnqh/QK5dxmJukeG8V87jhrIH6RQVgAGHfSbUV9aUfYINOq6vpdX1IxAuH0SPZD6PvEZSC/yLKliVw++QdaM+LwU69koIcGoaKellen57hrITyMnHfVv44WRe017eaYPOYJrwIwHewVgx9gruhjIOzq+AJ4A08Ab5APnBfL5rV/RBpDK5SYwhdwf7429aHTRqHP5aGUkGCeyPjFywAKwiFQIXeydAF+ROu9xxvEkEon7zhXaob6kDSzJ+wAAAABJRU5ErkJggg=="
        />
      </button>
    </div>
  </div>
</template>

<script>
import moment from "moment";

export default {
  name: "TaskItem",
  props: {
    task: Object,
  },
  data: () => {
    return {
      currentText: null,
      isWait: false,
      isEditable: false,
    };
  },
  created() {
    this.currentText = this.task.Text;
  },
  methods: {
    dateFilter: function (date) {
      if (date) return moment(date).fromNow();
      return "";
    },
    doneItem: function (task) {
      this.isWait = true;
      task.Done = !task.Done;
      this.$store.dispatch("updateTask", task).then(() => {
        this.isWait = false;
      });
    },
    updateText: function (task) {
      this.isWait = true;
      task.Text = this.currentText;
      this.$store.dispatch("updateTask", task).then(() => {
        this.isWait = false;
        this.isEditable = false;
      });
    },
    removeItem: function (task) {
      this.isWait = true;
      this.$store.dispatch("removeTask", task).then(() => {
        this.isWait = false;
      });
    },
    openEditable: function () {
      this.isEditable = true;
    },
  },
};
</script>

<style>
.task-item {
  position: relative;
  border: 1px solid;
  margin: 10px 0;
  padding: 15px 15px 15px 10px;
  font-size: 1.2em;
  line-height: 1;
  display: flex;
  align-items: center;
  outline: 1px solid;
}
.task-item.done {
  opacity: 0.5;
  outline: none;
}
.task-item.done span {
  font-style: oblique;
  text-decoration: line-through;
}

.task-item .editable {
  display: flex;
  width: 100%;
}
.task-item .editable input {
  font-size: 1.2em;
  width: 100%;
}
.task-item .editable button {
  margin-left: 3px;
}

.task-item small.date {
  min-width: 135px;
  margin-left: auto;
  opacity: 0.5;
  text-align: right;
}
.task-item:hover > small.date {
  visibility: hidden;
}

.task-item button {
  cursor: pointer;
  height: 32px;
  display: flex;
  padding: 0;
}
.task-item button img {
  object-fit: contain;
  height: 100%;
}
.task-item button.btn-done {
  margin-right: 10px;
}
.task-item .task-actions {
  position: absolute;
  top: 50%;
  transform: translateY(-50%);
  right: 10px;
  display: none;
}
.task-item .task-actions button {
  margin: 0 2px;
}
.task-item:hover > .task-actions {
  display: flex;
}
</style>