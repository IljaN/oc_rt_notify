<?php

namespace OCA\RealTimeNotifications;



use OCP\Util;

class ShareHooks {

	/** @var  EventRouterService */
	private static $er;


	public static function initialize(EventRouterService $er) {
		self::$er = $er;
		Util::connectHook('OCP\Share', 'post_shared', __CLASS__, 'itemShared');
	}

	/**
	 * @return mixed
	 */
	public static function getEventRouter() {
		return self::$er;
	}




	public static function itemShared(array $params) {
		self::$er->onPostShare($params);
	}

}