<!DOCTYPE html>
<html lang="zh-CN">
<head>
<meta charset="utf-8">
<meta http-equiv="X-UA-Compatible" content="IE=edge">
<meta name="viewport" content="width=device-width, initial-scale=1">
<!-- 上述3个meta标签*必须*放在最前面，任何其他内容都*必须*跟随其后！ -->
<title>输电线路远程巡视系统</title>

<!-- Bootstrap -->
<link rel="stylesheet" href="/static/css/bootstrap.min.css" >
<link rel="stylesheet" href="/static/css/bootstrap-treeview.min.css">
<link rel="stylesheet" href="/static/css/self.css">
<link rel="stylesheet" href="/static/css/dialog-polyfill.css">

<link rel="stylesheet" href="/static/css/material.css">
<!-- 
<link rel="stylesheet" href="https://fonts.googleapis.com/icon?family=Material+Icons">
-->
<link rel="import" href="photo_setting.html"> 

<script src="/static/js/dialog-polyfill.js"></script>
<script src="/static/js/material.js"></script>
<script src="/static/js/jquery-3.1.0.min.js"></script>
<script src="/static/js/bootstrap.min.js"></script>
<script src="/static/js/bootstrap-treeview.min.js"></script> 



<script type="text/javascript" src="http://api.map.baidu.com/api?v=2.0&ak=eE2DKK8r5Z9qVBs00FsDcRG3Mv3I6IGe"></script>

<!-- HTML5 shim and Respond.js for IE8 support of HTML5 elements and media queries -->
<!-- WARNING: Respond.js doesn't work if you view the page via file:// -->
<!--[if lt IE 9]>
                                            <script src="//cdn.bootcss.com/html5shiv/3.7.2/html5shiv.min.js"></script>
                                                  <script src="//cdn.bootcss.com/respond.js/1.4.2/respond.min.js"></script>
                                                      <![endif]-->
</head>
<body>
<div class="container-fluid c-f">
    <div class="row headerbgc"  >
        <div class="col-md-4">
            <button class="btn  btn-large btn-default button_style headerfont" type="button">输电线路远程巡视系统</button>
        </div>
        <div class="col-md-4">

                <!--


            <button type="button" class="btn btn-default">Default</button>
            <button type="button" class="btn btn-default">Default</button>
            <button type="button" class="btn btn-default">Default</button>
            <button type="button" class="btn btn-default">Default</button>

-->



        </div>
        <div class="col-md-4">

        </div>
    </div>
    <div class="row titlebg ">
        <div class="col-md-4">
            <button class="btn  btn-large btn-default button_style" type="button">当前用户：济南输电工区</button>
            <button type="button" class="btn btn-default">切换用户</button>
        </div>
        <div class="col-md-4">

        </div>
        <div class="col-md-4 ">
            <p id="timeid">当前时间：12:00</p>
        </div>
    </div>

    <div class="row baidumap">
        <div class="col-md-3">
            <ul class="nav nav-tabs">
                <li class="active"> <a href="#">设备</a> </li>
                <li> <a href="#">功能</a> </li>
            </ul>
            <div id='tree'></div>
        </div>
        <div class="col-md-9 baidumap" >
            <ul class="nav nav-tabs">
                <li id="device-map" class="active"> <a href="#">设备位置</a> </li>
                <li id="device-status"> <a href="#">设备状态</a> </li>
                <li id="device-info"> <a href="#">设备参数</a> </li>
            </ul>
            <div id="baidu-device-map" class="baidumap">
            </div>
            <div id="device-status" style="display: none;">
                <div class="row">
                    <div class="col-xs-6">
                        <div class="col-md-12"> 当前拍照尺寸 </div>

                        <button id="online" class="mdl-button mdl-js-button mdl-button--raised mdl-js-ripple-effect mdl-button--accent">
                            离线 
                        </button>
                        <div class="col-md-6"><label id="picture-date0" >拍照时间 2016－07-31 08:00:00 </label>
                        </div>
                        <div class="col-md-6">
                            MEID <button type="button" class="btn btn-default">MEIZ</button>
                        </div>
                        <div class="col-md-6">
                            电话号码 <button type="button" class="btn btn-default">123456887478</button>
                        </div>
                        <div class="col-md-6">
                            IMSI  <button type="button" class="btn btn-default">imsi</button>
                        </div>

                        <div class="col-md-12">
                            <img id="picture-0" src="../../../Public/1.jpg" width="100%" height="100%" />
                        </div>
                    </div>

                    <div class="col-xs-6">
                        <div class="col-md-12"> 设置历史图片 </div>
                        <div class="col-md-12">

                            <button id="device-setting-dialog" type="button" class="mdl-button mdl-button--raised mdl-js-button dialog-button" data-upgraded=",MaterialButton">拍照设置</button>
                            <button id="device-take-photo-dialog" type="button" class="mdl-button mdl-button--raised mdl-js-button dialog-button" data-upgraded=",MaterialButton">主动拍照</button>
                            <button type="button" class="btn btn-default">历史图片查询</button>
                            <button type="button" class="btn btn-default">微信推送</button>
                            <button type="button" class="btn btn-default">画面对比</button>
                            <button type="button" class="btn btn-default">视频拍照</button>
                            <input type="checkbox" name="" value="" />固定隐藏点
                        </div>

                        <div class="col-md-6">
                            <div ><label id="picture-date1" >拍照时间 2016－07-31 08:00:00 </label></div>
                            <img id="picture-1" src="../../../Public/1.jpg" width="100%" height="100％" />
                        </div> 

                        <div class="col-md-6">
                            <div ><label id="picture-date2" >拍照时间 2016－07-31 08:00:00 </label></div>
                            <img id="picture-2" src="../../../Public/1.jpg" width="100%" height="100％" />
                        </div> 
                        <div class="col-md-6">
                            <div ><label id="picture-date3" >拍照时间 2016－07-31 08:00:00 </label></div>
                            <img id="picture-3" src="../../../Public/1.jpg" width="100%" height="100％" />
                        </div> 
                        <div class="col-md-6">
                            <div ><label id="picture-date4" >拍照时间 2016－07-31 08:00:00 </label></div>
                            <img id="picture-4" src="../../../Public/1.jpg" width="100%" height="100％" />
                        </div> 
                        <div class="col-md-12"> 巡视状态 </div>
                        <div class="col-md-12">
                            当前状态 ： 正常     电量／时间： 6.8 v  ／  2016-07-06  07:20：09 
                        </div>
                        <div class="col-md-12">
                            隐患类型 ：<select>
                                <option value ="volvo" selected="selected"> 无缺陷</option>
                                <option value ="saab">Saab</option>
                                <option value="opel">Opel</option>
                                <option value="audi">Audi</option>
                            </select>
                            <input type="text" value="保存" size="5">
                            <input type="checkbox" name="" value="" />重点标记
                        </div>
                    </div>
                    <div id="device-info" style="display: none;">
                        设备信息
                    </div>
                </div>
            </div>

            <div class="row">
            </div>
        </div>



        <!-- setting dialog html start -->
        <dialog class="mdl-dialog" id="setting-dialog">
        <h6 class="mdl-dialog__title">设备设置</h6>
        <div class="mdl-dialog__content">
            <div class="row">
                <div class="col-md-12">
                    设备名称 : <button id="device" type="button" class="btn btn-default">设备名称</button>
                </div>
                <div class="col-md-6">
                    开机时间：
                    <div class="mdl-textfield mdl-js-textfield mdl-textfield--floating-label">
                        <input class="mdl-textfield__input" type="text" id="device-boot" >
                        <label class="mdl-textfield__label" for="sample1"></label>
                    </div>

                </div>
                <div class="col-md-6">
                    关机时间：
                    <div class="mdl-textfield mdl-js-textfield mdl-textfield--floating-label">
                        <input class="mdl-textfield__input" type="text" id="device-off" >
                        <label class="mdl-textfield__label" for="sample1"></label>
                    </div>

                </div>
                <div class="col-md-6">
                    开始工作时间：
                    <div class="mdl-textfield mdl-js-textfield mdl-textfield--floating-label">
                        <input   class="mdl-textfield__input" type="text" id="device-start-work" >
                        <label class="mdl-textfield__label" for="sample1"></label>
                    </div>

                </div>
                <div class="col-md-6">
                    照片像素：
                    <div class="mdl-textfield mdl-js-textfield mdl-textfield--floating-label">
                        <input class="mdl-textfield__input" type="text" id="" >
                        <label class="mdl-textfield__label" for="sample1">1500</label>
                    </div>
                    (万像素)  
                </div>
                <div class="col-md-6">
                    拍照次数：
                    <div class="mdl-textfield mdl-js-textfield mdl-textfield--floating-label">
                        <input class="mdl-textfield__input" type="text" id="sample1" >
                        <label class="mdl-textfield__label" for="sample1">8</label>
                    </div>
                    (次/天)  
                </div>
                <div class="col-md-6">
                    拍照间隔：
                    <div class="mdl-textfield mdl-js-textfield mdl-textfield--floating-label">
                        <input class="mdl-textfield__input" type="text" id="photo-fre" >
                        <label class="mdl-textfield__label" for="sample1">100</label>
                    </div>
                    (分钟)  
                </div>
                <div class="col-md-6">
                    摄像长度：
                    <div class="mdl-textfield mdl-js-textfield mdl-textfield--floating-label">
                        <input class="mdl-textfield__input" type="text" id="sample1" >
                        <label class="mdl-textfield__label" for="sample1">15</label>
                    </div>
                    (秒)  
                </div>
                <div class="col-md-6">
                    最大摄像次数：
                    <div class="mdl-textfield mdl-js-textfield mdl-textfield--floating-label">
                        <input class="mdl-textfield__input" type="text" id="sample1" >
                        <label class="mdl-textfield__label" for="sample1">1</label>
                    </div>
                    (次/天)  
                </div>
            </div>

        </div>
        <div class="mdl-dialog__actions">
            <button type="button" class="mdl-button" id="apply-this">应用到该设备</button>
            <button type="button" class="mdl-button" id="apply-all-line">应用到该路线下所有设备</button>
            <button type="button" class="mdl-button" id="apply-reboot">重启设备</button>
            <button type="button" class="mdl-button" id="apply-close">保存关闭</button>
        </div>
 
        </dialog>

        <!-- dialog html end -->

        <!-- setting dialog html start -->
        <dialog class="mdl-dialog" id="take-photo-dialog">
        <h6 class="mdl-dialog__title">实时拍照</h6>
        <div class="mdl-dialog__content">
            <!-- MDL Spinner Component -->
            <div class="mdl-spinner mdl-js-spinner is-active" id="mdl-spinner"></div>

        </div>
        <div class="mdl-dialog__actions">
            <button type="button" class="mdl-button" id="apply-close">关闭</button>
        </div> 
        </dialog>
        <!-- dialog html end -->

    <!-- jQuery (necessary for Bootstrap's JavaScript plugins) -->
    <!-- Include all compiled plugins (below), or include individual files as needed -->
    <script src="/static/js/self.js"></script> 
    </body>
    </html>


