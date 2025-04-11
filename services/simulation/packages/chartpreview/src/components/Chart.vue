<script setup lang="ts">

import {ColorType, createChart, CandlestickSeries, type UTCTimestamp} from "lightweight-charts";
import {onMounted, useTemplateRef} from "vue";

const chartElement = useTemplateRef("chartElement")

onMounted(() => {
  const chart = createChart(chartElement.value!, {
    width: 1000,
    height: 500,
    layout: {
      textColor: 'white',
      background: {
        type: ColorType.Solid,
        color: 'black',
      },
    }
  });

  const series = chart.addSeries(CandlestickSeries);
  series.setData([
    { open: 10, high: 10.63, low: 9.49, close: 9.55, time: 1642427876 as UTCTimestamp },
    { open: 9.55, high: 10.30, low: 9.42, close: 9.94, time: 1642514276 as UTCTimestamp },
    { open: 9.94, high: 10.17, low: 9.92, close: 9.78, time: 1642600676 as UTCTimestamp },
    { open: 9.78, high: 10.59, low: 9.18, close: 9.51, time: 1642687076 as UTCTimestamp },
    { open: 9.51, high: 10.46, low: 9.10, close: 10.17, time: 1642773476 as UTCTimestamp },
    { open: 10.17, high: 10.96, low: 10.16, close: 10.47, time: 1642859876 as UTCTimestamp },
    { open: 10.47, high: 11.39, low: 10.40, close: 10.81, time: 1642946276 as UTCTimestamp },
    { open: 10.81, high: 11.60, low: 10.30, close: 10.75, time: 1643032676 as UTCTimestamp },
    { open: 10.75, high: 11.60, low: 10.49, close: 10.93, time: 1643119076 as UTCTimestamp },
    { open: 10.93, high: 11.53, low: 10.76, close: 10.96, time: 1643205476 as UTCTimestamp },
    { open: 10.96, high: 11.90, low: 10.80, close: 11.50, time: 1643291876 as UTCTimestamp },
    { open: 11.50, high: 12.00, low: 11.30, close: 11.80, time: 1643378276 as UTCTimestamp },
    { open: 11.80, high: 12.20, low: 11.70, close: 12.00, time: 1643464676 as UTCTimestamp },
    { open: 12.00, high: 12.50, low: 11.90, close: 12.30, time: 1643551076 as UTCTimestamp },
    { open: 12.30, high: 12.80, low: 12.10, close: 12.60, time: 1643637476 as UTCTimestamp },
    { open: 12.60, high: 13.00, low: 12.50, close: 12.90, time: 1643723876 as UTCTimestamp },
    { open: 12.90, high: 13.50, low: 12.70, close: 13.20, time: 1643810276 as UTCTimestamp },
    { open: 13.20, high: 13.70, low: 13.00, close: 13.50, time: 1643896676 as UTCTimestamp },
    { open: 13.50, high: 14.00, low: 13.30, close: 13.80, time: 1643983076 as UTCTimestamp },
    { open: 13.80, high: 14.20, low: 13.60, close: 14.00, time: 1644069476 as UTCTimestamp },
  ]);

  chart.timeScale().resetTimeScale();
  chart.timeScale().fitContent();


});


</script>

<template>
  <div ref="chartElement"></div>
</template>

<style scoped>

</style>
