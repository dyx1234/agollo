/*
 * Licensed to the Apache Software Foundation (ASF) under one or more
 * contributor license agreements.  See the NOTICE file distributed with
 * this work for additional information regarding copyright ownership.
 * The ASF licenses this file to You under the Apache License, Version 2.0
 * (the "License"); you may not use this file except in compliance with
 * the License.  You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package configmap

import (
	"github.com/apolloconfig/agollo/v4/component/log"
	"github.com/apolloconfig/agollo/v4/env/config"
)

type Store struct {
	K8sManager *K8sManager
}

// LoadConfigMap load ApolloConfig from configmap
func (c *Store) LoadConfigMap(appConfig config.AppConfig, configMapNamespace string) (*config.ApolloConfig, error) {
	var config = &config.ApolloConfig{}
	configMapName := appConfig.AppID
	key := appConfig.Cluster + "-" + appConfig.NamespaceName
	// TODO 在这里把json转为ApolloConfig, 但ReleaseKey字段会丢失, 影响大不大
	config.Configurations, _ = c.K8sManager.GetConfigMap(configMapName, configMapNamespace, key)

	config.AppID = appConfig.AppID
	config.Cluster = appConfig.Cluster
	config.NamespaceName = appConfig.NamespaceName
	return config, nil
}

// WriteConfigMap write apollo config to configmap
func (c *Store) WriteConfigMap(config *config.ApolloConfig, configMapNamespace string) error {
	// AppId作为configMap的name,cluster-namespace作为key, value为config
	configMapName := config.AppID
	key := config.Cluster + "-" + config.NamespaceName
	err := c.K8sManager.SetConfigMap(configMapName, configMapNamespace, key, config)
	if err != nil {
		log.Errorf("Failed to write ConfigMap %s in namespace %s: %v", configMapName, configMapNamespace, err)
		return err
	}
	return nil
}
