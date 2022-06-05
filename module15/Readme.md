#云原生训练营总结
  
>经过了三个多月训练营的学习，自己对云原生的整体有了一个更加深入的理解，对Kubernetes 的运维管理 与开发都有了进一步的理解，感谢这几个月老师们的教导，也感谢自己能够坚持下来。下面我对每一章节的内容做一下自己的总结。  

>第一章讲解了Go语言的基本语法与特性，并较为深入的讲解了生产者与消费者模型的Go语言实现。Go语言与其他的语言相比，有更好的单点编译能力，减轻开发对于底层类型关系处理的复杂度，支持垃圾回收，有着更好的并发与多线程的能力。Go 语言有着对开发更加友好的能力，自身把一些复杂问题简单化，从而降低了开发编程的复杂度，提升了开发的效率。

>第二章讲解了Go语言的线程加锁，线程调度，内存管理，包引入与依赖管理。因为计算机的为了提升计算的速度，加了各种的缓冲，比如CPU 的Cache,内存,与磁盘缓存等等，CPU 计算时候一般都是从内存中拿到值到CPU 中计算，而计算的中间值一般存在与CPU的缓存中，在多个线程多CPU同时操作一个值得情况下就会产生不可控的情况。因此我们需要对一些多线程共享的值，进行加锁来保证线程的安全性。进程是资源分配的基本单位，线程是调度的基本单位。进程切换的开销在于直接开销与间接开销，直接开销包括切换页表全局目录，切换内核动态栈，切换硬件上下文，刷新TLB,系统调度器的代码执行，间接开销为CPU缓存失效导致的进程需要到内存访问，增加IO操作。线程的本质上只是一批共享资源的进程，线程切换本质上依然需要内核进行进程切换，一组线程因为共享内存资源，因此一个进程的所有线程共享虚拟地址空间，线程切换相比进程切换，主要节省了虚拟地址空间的切换。

>第三章讲解了Docker 的核心技术。Docker 基于Linux 内核的Cgroup,Namespace，以及Union FS 等技术，对进程进行封装隔离，属于操作系统层面的虚拟化技术，由于隔离的进程独立于宿主和其他的隔离的进程，因此也称其为容器。最初实现是基于LXC,后面转而自行开发Libcontainer ，后面进一步演化为runC 和Containerd。Docker 在容器的基础上，进一步的封装，从文件系统，网络互联到进程隔离等等，极大的简化了容器的创建和维护，使得Docker 技术比虚拟机技术更为轻便、快捷。Docker 技术能更高效的利用系统资源，更快速的启动时间，一致的运行环境，持续交付和部署，更加轻松的迁移与维护扩展。容器主要特性有安全性，隔离性，便携性，可配额。隔离性，Linux Namespace 实现隔离，Namespace 的类型有IPC,Network,PID，Mount,UTS,USR等。可配额主要有Cgroups 技术实现。Cgroups 实现了对blkio,cpu,cpuacct,cpuset,devices,freezer,memory,net_cls,ns,pid资源的配置和度量。

>第四章讲解了Kubernetes 的架构原则与对象设计。Kubernetes 是谷歌开源的容器集群管理系统，是Google多年大规模容器管理技术Borg 的开源版本，主要功能包括：基于容器的应用部署、维护和滚动升级你；负载均衡和服务发现；跨机器和跨地域的集群调度；自动伸缩；无状态服务和有状态服务；插件机制保证扩展性。Kubernetes 是声明式系统，核心对象包括：抽象计算节点Node,隔离资源的基本单位Namespace,描述应用实例的Pod ，服务发现Service。核心组件有Mster Node: API服务器API Server， 集群数据存储Etcd,控制管理器Controller Manager, 调度器Scheduler；Work Node: Pod 生命周期管理Kubelet, 网络转发与负载均衡Kube-proxy。

>第五章讲解了Kubernetes 控制平面组件Etcd。Etcd 是CoreOS基于Raft开发的分布式key-value存储，可用于服务发现，共享配置以及一致性保障。他提供了对数据TTL失效，数据改变监视，多值，目录监听，分布式锁原子操作等功能，可能方便的跟踪并管理集群节点的状态。Etcd基于Raft协议开发的，Raft 协议基于quorum机制，即大多数同意原则，任何的变更都需要超过半数的成员确认。Etcd 基于Raft的一致性，选举方法： 初始启动时，节点处于Follower 状态并被设定一个election timeout ，如果在这一时间周期内没有收到来自Leader的heartbeat， 节点将发起选举： 将自己切换为candidate 之后，向集群中其他Follower 节点发送请求，询问其是否选举自己成为Leader。 当收到来自集群中过半节点的接受投票后，节点即成为Leader，开始接受保存client 的数据并向其他的Follower 节点同步日志。如果没有达成一致，则candidate 随机选择一个等待间隔（150ms  300ms) 在发起投票，投票中过半的Follower 将成为Leader。Leader 节点依靠Heartbeat 保存地位，新Leader 比老Leader 任期大1。

>第六章讲解了kubernetes 的核心组件之一kube-apiserver。kube-apiserver 包含了认证，鉴权，准入（Mutating,Validating,Admission) ，限流，API Server 对象的实现。提供其他模块之间的数据交互和通信枢纽(其他模块通过API Server 查询或修改数据，只有API Server才直接操作etcd)。

>第七章讲解了Kubernetes 控制平面组件kube-scheduler ，controller， kubelet等。 Kube-scheduler 负责分配调度Pod 到集群内的节点，它监听kube-apiserver,查询还没有分配的Node 的Pod, 然后跟读调度策略把Pod 分配节点(更新Pod 的NodeName 字段)。调度器需要充分考虑的因素：公平调度，资源高效利用，Qos, affinity 和 anti-affinity ，数据本地化，内部负载干扰。deadlines。Controller 维护集群各个资源的生命周期，对Node,Pod 等资源状态的更新与维护。Kubelet 负责Pod 的生命周期的监控与维护，集成了CRI ,CNI,CSI 的接口方便对接各种容器，网络，存储等相关不同的解决方案。

>第八章讲解了控制平面组件生命周期管理和服务发现。深入理解Pod 的生命周期：用户通过命令发出创建pod 的命令（Pod 生命周期出于Create），命令到达APISERVER,等待Scheduler 的调度(Pod 生命周期处于Peding)，调度器调度到可用node 节点上，kubelet 负载拉起服务(Pod 生命周期处于containerCreating)。kubelet 拉起pod 正常运行（Pod 生命周期出于Runing）。客户端发出Delete 命令（Pod 生命周期处于Terminating）,删除完成（Pod 生命周期处于gone）。如果kubelet 因为网络问题等原因无法获取Pod的状态（Pod 生命周期处于Failed 或者Unknown）, 如果Kubelet 被驱逐（Pod 生命周期处于Succeeded,Evicted）。服务发现，需要把服务发布至集群内部或者外部，服务的不同类型： ClusterIP(Headless)、NodePort、LoadBalancer、ExternalName。

>第九章讲解了Kubernetes 生产集群的管理。云原生的原则，可变基础设施是存在风险的，因为持续手工操作会引入很多不可控的情况，导致不可预知的问题，因此推荐使用不可变的容器镜像，不可变的操作系统。操作系统的原则是最小化主机操作系统，只安装必要的工具，这样的意义在于有更好的性能，稳定性，安全保障。构建高可用Kuberentes 集群保证Controller panle 的可用性，从物理层面保证不因电源导致局部断掉，整个控制平面失效。部署多个控制平面让使用负载均衡器提供给WorkNode 服务。

>第十章讲解了Kubernertes 的生产化运维。Kubernetes 的生产化运维包括镜像仓库，镜像的安全，基于Kubernetes 的DevOps ,基于GitHuB Action 的自动化流水线，基于Jenkins 的自动化流水线，云原生流水线Tekton,持续部署ArgoCD ，监控Prometheus stack,与日志Loki stack。

>第十一章讲解了将应用迁移至Kubernetres 平台。 应用迁移至Kubernetes 需要做： 应用容器化，区分有状态和无状态使用不同的方式进行迁移。使用应用管理器Helm 有助于管理迁移后的App， 部署metrics-server 让集群能够使用HPA 自动扩缩容能力。节点自动扩缩容使用VPA。

第十二章讲解了基于Istio 的高级流量管理。Istio 具有适应性（熔断，重试，超市处理，失败处理，负载均衡，Failover），服务发现（路由），安全和访问控制（TLS和证书管理），可观察行（Metrics 监控 分布式日志 分布式 tracing）。Istio 的流量管理是通过sidecar 进行劫持流量。

>第十三章讲解了Kubernetes 集群联邦和Istio 多集群管理。Kuberntes 集群联邦是对多集群的治理能力，能够进行成本优化，避免厂商绑定，容灾，提升响应速度等等。Istio 在多集群管理中，使用权重进行流量的控制，是的多数流量转发到本地域，少数流量转发到跨地域，在本地域发生故障时，能够很好的进行容灾处理。

>第十四章讲解了在云原生下保障容器运行的安全。传统的安全三元素CIA,机密性（Confidentiality）,完整性（Integrity），可用性（Availability）。对于安全需要从工具到方法上实现。容器以Non-root 的身份运行，集群的安全性一般常用手段是Pod 安全上下文，API Server 的认证，授权，审计和准入控制，数据的加密机制等。集群的通讯安全，默认情况下API 的通信都使用TLS 加密。如果你希望在IP地址或端口层面（OSI 第3层或第4层）控制网络流浪，那么你可以考虑为集群中特性应用使用kubernetes 网络策略（NetworkPolicy）,Pod 可以通讯的Pod 是通过如下三个标识符的组合来辨识的：其他被允许的Pod;被允许的命名空间；IP组块。

>第十五章讲解了微服务的开发和部署案例。Kubernetes 构建多租户平台：基于Tenant 的租户管理，基于Application 的应用治理。Tenant 是管理应用的账号，Tenant 也是集群费用分摊实体，Tenant 是多租户集群中的租户。Bookinfo 的应用架构解析微服务在Kuberentes 上的治理。

>虽然在这三个多月的学习中，还有很多需要继续提升的，但是能坚持到现在也是对自己的一种肯定。
