package main

import (
	"flag"
	"os"

	"github.com/go-logr/zapr"
	"go.uber.org/zap"
	ctrl "sigs.k8s.io/controller-runtime"
)

func main() {
	var metricsAddr string
	var enableLeaderElection bool

	flag.StringVar(&metricsAddr, "metrics-bind-address", ":8080", "The address the metric endpoint binds to.")
	flag.BoolVar(&enableLeaderElection, "leader-elect", false, "Enable leader election for controller manager.")
	flag.Parse()

	zapLog, _ := zap.NewProduction()
	ctrl.SetLogger(zapr.NewLogger(zapLog))

	setupLog := ctrl.Log.WithName("setup")
	setupLog.Info("Starting K8s Multi-Cluster Orchestrator Manager v1.0.0")

	mgr, err := ctrl.NewManager(ctrl.GetConfigOrDie(), ctrl.Options{
		LeaderElection: enableLeaderElection,
	})
	if err != nil {
		setupLog.Error(err, "unable to start manager")
		os.Exit(1)
	}

	setupLog.Info("Controller manager initialized successfully")
	if err := mgr.Start(ctrl.SetupSignalHandler()); err != nil {
		setupLog.Error(err, "problem running manager")
		os.Exit(1)
	}
}
