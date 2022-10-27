package web

import (
	"ngrok/log"
	"time"

	"github.com/jmoiron/sqlx"
	"github.com/shirou/gopsutil/cpu"
	"github.com/shirou/gopsutil/disk"
	"github.com/shirou/gopsutil/load"
	"github.com/shirou/gopsutil/mem"
	"github.com/shirou/gopsutil/net"
	"github.com/shirou/gopsutil/process"
)

func readInfo(db *sqlx.DB) {
	t_net_i := float64(0)
	t_net_o := float64(0)
	t_disk_r := float64(0)
	t_disk_w := float64(0)
	for {
		percent, _ := cpu.Percent(time.Second, false)
		memInfo, _ := mem.VirtualMemory()
		d, _ := disk.IOCounters()
		nv, _ := net.IOCounters(false)
		ld, _ := load.Avg()
		pd, _ := process.Pids()

		net_i := float64(0)
		net_o := float64(0)
		for _, n := range nv {
			if n.BytesRecv == 0 && n.BytesSent == 0 {
				continue
			}
			net_i += float64(n.BytesRecv)
			net_o += float64(n.BytesSent)
		}
		tmp_i := net_i
		tmp_o := net_o
		if t_net_i != 0 || t_net_o != 0 {
			net_i -= t_net_i
			net_o -= t_net_o
		} else {
			net_i = float64(0)
			net_o = float64(0)
		}
		t_net_i = tmp_i
		t_net_o = tmp_o

		disk_r := float64(0)
		disk_w := float64(0)
		for _, v := range d {
			if v.ReadBytes == 0 && v.WriteBytes == 0 {
				continue
			}
			disk_r += float64(v.ReadBytes)
			disk_w += float64(v.WriteBytes)
		}

		tmp_r := disk_r
		tmp_w := disk_w
		if t_disk_r != 0 || t_disk_w != 0 {
			disk_r -= t_disk_r
			disk_w -= t_disk_w
		} else {
			disk_r = float64(0)
			disk_w = float64(0)
		}
		t_disk_r = tmp_r
		t_disk_w = tmp_w

		sys := &SystemInfo{
			Cpu:        percent[0],
			Mem:        memInfo.UsedPercent,
			DiskR:      disk_r,
			DiskW:      disk_w,
			NetI:       net_i,
			NetO:       net_o,
			Load:       ld.Load1,
			Pid:        float64(len(pd)),
			CreateTime: time.Now(),
		}

		_, err := db.NamedExec("insert into `system_info`(`cpu`,`mem`,`disk_r`, `disk_w`,`net_i`,`net_o`,`load`,`pid`,`create_time`) values(:cpu,:mem,:disk_r,:disk_w,:net_i,:net_o,:load,:pid,:create_time)", sys)
		if err != nil {
			log.Info("%v", err.Error())
		}
		sys.CreateTime = sys.CreateTime.Add(-240 * time.Hour)
		db.NamedExec("delete from `system_info` where `create_time` <= :create_time", sys)
	}
}
